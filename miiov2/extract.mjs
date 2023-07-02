import * as fs from 'fs';

const devices = JSON.parse(fs.readFileSync('./specs.json'));

const services = new Map();
const properties = new Map();
const actions = new Map();

devices.forEach(device => {
  device.services.forEach(service => {
    const type = shortType(service.type);
    if (!services.has(type)) {
      services.set(type, []);
    }
    services.get(type).push(service);
    service.properties?.forEach(property => {
      const type = shortType(property.type);
      if (!properties.has(type)) {
        properties.set(type, []);
      }
      properties.get(type).push(property);
    });
  });
});

const devicesGenFile = fs.createWriteStream('./devices/gen.go');
devicesGenFile.write('package devices\n');
// devicesGenFile.write('import "sup/miio/properties"\n');
devices.forEach(device => {
  const { name: deviceType, device: deviceName } = parseURN(device.type);
  const name = camelCase(`${deviceType}-${deviceName}`);
  const res = [
    `// ${name} ${device.description}`,
    `var ${name} = struct {`,
  ];
  device.services.forEach(service => {
    const serviceName = service.description.replace(/ Service$/i, '').replace(/ /g, '');
    service.properties?.forEach(property => {
      const propName = shortType(property.type, false);
      res.push(`  ${camelCase(`${serviceName}-${propName}`)} PropertyBinding`);
    });
  });
  res.push('} {')
  device.services.forEach(service => {
    const serviceName = service.description.replace(/ Service$/i, '').replace(/ /g, '');
    service.properties?.forEach(property => {
      const propName = shortType(property.type, false);
      res.push(`  ${camelCase(`${serviceName}-${propName}`)}: PropertyBinding{${service.iid}, ${property.iid}},`);
    });
  });
  res.push('}');
  devicesGenFile.write('\n');
  devicesGenFile.write(res.join('\n'));
});
devicesGenFile.end();

const servicesGenFile = fs.createWriteStream('./services/gen.go');
servicesGenFile.write('package services\n');
servicesGenFile.write('import "sup/miio/properties"\n');
services.forEach(services => {
  const s = new Set(services.map(service => {
    const name = camelCase(shortType(service.type));
    const res = [
      `// ${name} ${service.description}`,
      `type ${name} struct {`
    ];
    service.properties?.forEach(property => {
      const propName = camelCase(shortType(property.type, false));
      const name = camelCase(shortType(property.type));
      res.push(`  ${propName} properties.${name}`);
    });
    res.push('}');
    return res.join('\n');
  }));
  // console.log(s);
  if (s.size > 1) {
    console.warn('multiple services with same type', s);
  }
  servicesGenFile.write('\n');
  servicesGenFile.write(s.values().next().value);
});
servicesGenFile.end();

const propertiesGenFile = fs.createWriteStream('./properties/gen.go');
propertiesGenFile.write('package properties\n');
properties.forEach(properties => {
  const s = new Set(properties.map(property => {
    const name = camelCase(shortType(property.type));
    const res = [
      `// ${name} ${property.description}`,
      `type ${name} ${toGoType(property.format)}`,
    ];
    const enums = property['value-list'];
    if (enums) {
      res.push(`const (`);
      enums.forEach(({ value, description }) => {
        res.push(`  ${name}${camelCase(description)} ${name} = ${value}`);
      });
      res.push(`)`);
    }
    // res.push(`func (p *${name}) Unit() Unit { return Unit${camelCase(property.unit || 'none')} }`);
    return res.join('\n');
  }));
  // s.forEach(s => console.log(s));
  if (s.size > 1) {
    console.warn('multiple properties with same type', s);
  }
  propertiesGenFile.write('\n');
  propertiesGenFile.write(s.values().next().value);
});
propertiesGenFile.end();

function parseURN(urn) {
  const [, spec, ns, name, id, device, deviceVersion] = urn.split(/:/g)
  return { spec, ns, name, id, device, deviceVersion };
}

function shortType(urn, withSpec = true) {
  const { spec, name } = parseURN(urn);
  if (withSpec) {
    return spec.replace(/-spec/, '') + '-' + name;
  }
  return name;
}

function camelCase(str) {
  return str.replace(/[^a-z0-9]*\b([a-z])/ig, (m, c) => c.toUpperCase());
}

function lowerCamelCase(str) {
  return camelCase(str).replace(/^[A-Z]/, c => c.toLowerCase());
}

function toGoType(format) {
  switch (format) {
    case 'float':
      return 'float64';
    default:
      return format;
  }
}
