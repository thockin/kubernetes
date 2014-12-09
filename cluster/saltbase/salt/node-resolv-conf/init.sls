{% set dns_server=pillar['dns_server'] %}
{% set dns_domain=pillar['dns_domain'] %}

/etc/resolv.conf-nameserver:
  file.blockreplace:
    - name: /etc/resolv.conf
    - marker_start: '#-- start kubernetes managed section: nameserver --'
    - marker_end: '#-- end kubernetes managed section: nameserver --'
    - content: 'nameserver {{dns_server}}'
    - prepend_if_not_found: True

/etc/resolv.conf-search:
  file.replace:
    - name: /etc/resolv.conf
    - pattern: 'search *(.*)'
    - repl: |
        search \1
        #-- start kubernetes managed section: search --
        search {{dns_domain}}. \1
        #-- end kubernetes managed section: search --
