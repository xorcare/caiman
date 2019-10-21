# caiman

This is a tool for converting [LDIF] ([LDAP]) data into standard [vCard4] data, and
then downloading it to services or applications that are compatible with the
vCard4 standard.

The tool is mainly focused on the [Microsoft AD] product, but can be used with other
LDAPs, with possible limitations.

## Command line interface, [CLI]

```
caiman is a tool to convert people data from LDAP(LDIF) format to vCard4 contact format

Usage:
  caiman [flags]

Examples:
caiman < person.ldif > person.vcf
caiman --config-file ~/.caiman.yaml < person.ldif > person.vcf
caiman --config-dump > .caiman.yaml
cat person.ldif | caiman > person.vcf
cat person.ldif | caiman | tee person.vcf

Flags:
  -d, --config-dump          print to standard output all configuration values, it prints configuration data in YAML format
  -f, --config-file string   the settings file from which the settings will be loaded
  -h, --help                 help for caiman
```

Example of a command to convert LDAP to vCard:

```
$ caiman < person.ldif > person.vcf 
2019/01/31 21:01:44 total entries 1030
2019/01/31 21:01:44 skipped 1 entries because it is nil
2019/01/31 21:01:44 skipped 5 entries because bad count of filled fields
2019/01/31 21:01:44 successfully processed 1024 entries
```

Example of a command to export data from LDAP:

```
ldapsearch -x -h example.com -D "DOMAIN\user.name" -LL -W -b 'DC=example,DC=com' '(objectClass=person)' | tee person.ldif
```

## License

© Vasiliy Vasilyuk, 2019-2020

Released under the [BSD 3-Clause License][LIC].

[LIC]:https://github.com/xorcare/caiman/blob/master/LICENSE 'BSD 3-Clause "New" or "Revised" License'
[vCard4]:https://en.wikipedia.org/wiki/VCard#vCard_4.0 'vCard 4.0 its the latest standard, which is built upon the RFC 6350 standard'
[LDIF]:https://tools.ietf.org/html/rfc2849 'The LDAP Data Interchange Format (LDIF) - Technical Specification'
[LDAP]:https://en.wikipedia.org/wiki/LDAP_Data_Interchange_Format 'LDAP Data Interchange Format'
[Microsoft AD]:https://docs.microsoft.com/en-us/azure/active-directory 'Azure Active Directory documentation'
[CLI]:https://en.wikipedia.org/wiki/Command-line_interface 'Command-line interface'
