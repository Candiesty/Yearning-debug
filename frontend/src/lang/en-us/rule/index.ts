export default {
  DMLTransaction: 'DML statements are executed using a transaction',
  DDLCheckTableComment: 'Tables must have table comment',
  DDlCheckColumnComment: 'Table fields must have column comments',
  DDLCheckColumnNullable: 'Fields of non-TIMESTAMP type must be NOT NULL',
  DDLCheckColumnDefault:
    'Non-text,blob, JSON, TIMESTAMP fields must have default values',
  DDLCheckFloatDouble: 'Force the float/double type to be of type Decimal',
  DDLEnableAutoincrementInit: 'The increment column to start with 1',
  DDLPrimaryKeyMust: 'Force the primary key name to be ID',
  DDLEnableAutoIncrement: 'Forces primary keys to increment columns',
  DDLEnableAutoincrementUnsigned:
    'The primary key must use the unsigned flag unsigned',
  DDLIndexNameSpec:
    'Index name specification (index names must begin with idx_)',
  CheckIdentifier: 'Enable Mysql keyword Check',
  DDLEnableAcrossDBRename: 'Allow migration across tables',
  DDLEnableDropTable: 'Table deletion is allowed',
  DDLEnableDropDatabase: 'Allow database deletion',
  DDLAllowPRINotInt: 'Non-int /bigint primary key type is allowed',
  DDLEnableNullIndexName: 'Empty index name is allowed',
  DDLMultiToCommit:
    'Allows a single work order to submit multiple DDL statements',

  DDLAllowMultiAlter:
    'Allows a single work order to execute multiple ALTER statements',
  DDLAllowColumnType:
    'Allows fields to be typed (converted between different fields or changed from long to short in length). For example :int -> bigint,int(50) -> int(20))',
  DDLAllowChangeColumnPosition: 'After /first is allowed',
  AllowCreateView: 'Allow creation of views',
  AllowCreatePartition: 'Allow partition creation',
  AllowSpecialType: 'Bit,enum, and SET fields are allowed',
  SupportCollation:
    'The Collate range allowed when CREATE/ALTER a table or field. Use commas to separate multiple items',
  SupportCharset:
    'The RANGE OF CHARACTER sets ALLOWED WHEN CREATE/ALTER A TABLE or FIELD. Use commas to separate multiple items',
  MustHaveColumns:
    'Table must have fields. Separate multiple fields with commas',
  DDLMaxKeyParts: 'A single index specifies the upper limit of a field',
  DDLMaxKey: 'A single table allows a maximum of several indexes',
  MaxDDLAffectRows: 'Maximum number of rows affected by DDL',
  DDLMaxCharLength: 'Char Indicates the maximum length of the char field',
  MaxTableNameLen: 'Maximum length of a table name',
  DMLMaxInsertRows: 'Insert Indicates the maximum number of rows inserted',
  DMLAllowLimitSTMT:
    'The limit keyword is allowed for UPDATE/INSERT statements',
  DMLAllowInsertNull: 'Insert statements are allowed to insert null values',
  DDLImplicitTypeConversion: 'Implicit conversions are not allowed',
  DMLInsertColumns:
    'Check whether the field name inserted in the Insert statement exists',
  DMLWhere: 'Enforce that DML statements must have a WHERE condition',
  DMLOrder: 'Disallow DML statements from using the Order BY clause',
  DMLSelect: 'Disallow Select clauses for DML statements',
  MaxAffectRows: 'DML affects the maximum number of rows',
  IsOSC: 'Start the online table change tool',
  OscSize:
    'When the volume of a table exceeds the value, online synchronization of table changes is triggered. The unit is M',
  OSCExpr:
    'Synchronize tool parameters. For example: PT-OSC please note! Set only parameters here. Replace the input value $SQL $ADDR $PORT $USER $PASSWORD $SCHEMA $TABLE with the following variable name. Example: (PT-OSC configuration) pt-online-schema-change  --alter $SQL --user=$USER  --password=$PASSWORD  --host=$ADDR P=$PORT,D=$SCHEMA,t=$TABLE  --execute',
  global: 'Global rules',
  custom_list: 'Custom rule list',
  custom: 'Custom rule',
};
