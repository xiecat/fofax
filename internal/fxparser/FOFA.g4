grammar FOFA;

options {
    language = Go;
}

start: query;

query
   : leftBracket = BR_OPEN query rightBracket = BR_CLOSE                    #bracketExp
   | leftQuery=query op=AND rightQuery=query                                #andLogicalExp
   | leftQuery=query op=OR rightQuery=query                                 #orLogicalExp
   | propertyName=attrPath op=COMPARISON_OPERATOR propertyValue=attrValue   #compareExp
   | propertyName=attrPath op=SCOMPARISON_OPERATOR propertyValue=attrValue  #scompareExp
   ;

attrPath
   : ATTRNAME
   ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

attrValue
   : BOOLEAN           #boolean
   | NULL              #null
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #long
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

NULL
   : 'null'
   ;

COMPARISON_OPERATOR: EQ;
SCOMPARISON_OPERATOR: SEQ;

AND:   '&&';
OR:    '||';
EQ : '=';
SEQ: '==';
BR_OPEN: '(';
BR_CLOSE: ')';

ATTRNAME
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
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;
// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;
// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

WS: [\t ]+ -> skip;
