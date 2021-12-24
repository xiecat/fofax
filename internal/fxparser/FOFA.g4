grammar FOFA;

options {
    language = Go;
}

start: query;

query
   : leftBracket = BR_OPEN query rightBracket = BR_CLOSE                    #bracketExp
   | leftQuery=query op=AND rightQuery=query                                #andLogicalExp
   | leftQuery=query op=OR rightQuery=query                                 #orLogicalExp
   | propertyName=fofaKeyword op=EQ propertyValue=fofaValue                    #compareExp
   | propertyName=fofaKeyword op=SEQ propertyValue=fofaValue                   #scompareExp
   | propertyName=fofaKeyword op=NOT propertyValue=fofaValue                   #noCompareExp
   | sgatom=fofaValue              #sgExp
   ;

fofaKeyword
   : FOFA_KEY('.' FOFA_KEY)?;

fofaValue
   : BOOLEAN              #boolean
   | STRING               #string
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

BOOLEAN
   : 'true' | 'false'
   ;

AND:   '&&';
OR:    '||';
NOT: '!=';
EQ : '=';
SEQ: '==';
BR_OPEN: '(';
BR_CLOSE: ')';


FOFA_KEY
   : 'title'
   | 'header'
   | 'body'
   | 'fid'
   | 'domain'
   | 'icp'
   | 'js_name'
   | 'js_md5'
   | 'icon_hash'
   | 'host'
   | 'port'
   | 'ip'
   | 'status_code'
   | 'protocol'
   | 'country'
   | 'region'
   | 'city'
   | 'cert'
   | 'cert.subject'
   | 'cert.issuer'
   | 'cert.is_valid'
   | 'banner'
   | 'type'
   | 'os'
   | 'before'
   | 'after'
   | 'org'
   | 'base_protocol'
   | 'server'
   | 'app'
   | 'asn'
   | 'cname'
   | 'cname_domain'
   | 'is_fraud'
   | 'is_honeypot'
   | 'is_ipv6'
   | 'is_domain'
   | 'port_size'
   | 'port_size_gt'
   | 'port_size_lt'
   | 'ip_ports'
   | 'ip_country'
   | 'ip_region'
   | 'ip_city'
   | 'ip_after'
   | 'ip_before'
   | 'fx'
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   | ([a-zA-Z0-9.+_-]|('\\' ["\\/bfnrt]))*
   ;

WS: [\t ]+ -> skip;