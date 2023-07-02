import * as fs from 'fs';
import path from 'path';

const devices = JSON.parse(fs.readFileSync('./specs.json'));

devices.forEach(device => {
  const urn = parseURN(device.type);
  const properties = new Map();

  device.services.forEach(service => {
    service.properties?.forEach(property => {
      const propType = camelCase(shortType(property.type, false));
      if (!properties.has(propType)) {
        properties.set(propType, []);
      }
      properties.get(propType).push(property);
    });
  });

  const res = [
    `package ${packageName(urn.device)}`,
    `import (`,
    `  "reflect"`,
    `  "fmt"`,
    ``,
    `  "sup/miio/devices"`,
    `)`,
    `var _ fmt.Stringer`,
  ];

  res.push(`func init() {`);
  device.services.forEach(service => {
    const serviceName = service.description.replace(/ Service$/i, '').replace(/ /g, '');
    service.properties?.forEach(property => {
      const propName = camelCase(shortType(property.type, false));
      const name = camelCase(serviceName) + camelCase(propName);
      res.push(`  devices.RegisterProperty(reflect.TypeOf(${name}{}), ${service.iid}, ${property.iid})`);
    });
  });
  res.push(`}`);

  device.services.forEach(service => {
    const serviceName = service.description.replace(/ Service$/i, '').replace(/ /g, '');
    service.properties?.forEach(property => {
      const propName = camelCase(shortType(property.type, false));
      const propType = camelCase(shortType(property.type, false));
      const name = camelCase(serviceName) + camelCase(propName);
      res.push(`// ${name} ${property.description}`);
      res.push(`type ${name} struct {`);
      res.push(`  V ${propType}V`);
      res.push(`  Err error // Output only`);
      res.push(`}`);
    });
  });

  properties.forEach((properties, propType) => {
    const s = new Set();
    properties.forEach(property => {
      const t = [
        `// ${propType}V ${property.description}`,
        `type ${propType}V ${toGoType(property.format)}`,
      ];
      const enums = property['value-list'];
      if (enums) {
        t.push(`const (`);
        enums.forEach(({ value, description }) => {
          t.push(`  ${camelCase(propType)}${camelCase(description)} ${propType}V = ${value}`);
        });
        t.push(`)`);
        t.push(`func (v ${propType}V) String() string {`);
        t.push(`  switch v {`);
        enums.forEach(({ value, description }) => {
          t.push(`  case ${camelCase(propType)}${camelCase(description)}:`);
          t.push(`    return "${description}"`);
        });
        t.push(`  default:`);
        t.push(`    return fmt.Sprintf("unknown ${propType}: %v", ${toGoType(property.format)}(v))`);
        t.push(`  }`);
        t.push(`}`);
      }
      s.add(t.join('\n'));
    });
    if (s.size > 1) {
      console.log(`WARN: ${urn.device}.${propType} has multiple definitions`, s);
    }
    res.push(s.values().next().value);
  });

  const outputDir = path.join("devices", packageName(urn.name) + 's', packageName(urn.device));
  const outputFile = path.join(outputDir, "gen.go");
  fs.mkdirSync(outputDir, { recursive: true });
  fs.writeFileSync(outputFile, res.join('\n'));
});

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

function packageName(str) {
  return str.toLowerCase().replace(/[^a-z0-9]/ig, '');
}

function camelCase(str) {
  return str.replace(/[^a-z0-9]*\b([a-z])/ig, (m, c) => c.toUpperCase()).replace(/\W+/g, '');
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
