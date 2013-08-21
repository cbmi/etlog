# etlog

etlog is service for logging metadata during an ETL process. One of the challenges with complex ETL is keeping track of **where data comes from**, **how it was processed**, and **when it populated**. etlog messages contain information about the the _source_ and _target_ data as well as the _script_ that performed the ETL.

The intended outcomes of etlog is to enable:

- Linking production data back to the source data
- Capturing the necessary metadata to inform data changes over time

The features implemented to support the above outcomes include:

- Flexible document-based data model for storing as much or as little necessary to provide utility
- Built-in support for common data stores such as files, relational databases, document types, etc.

## Transaction

During an ETL process, an action is performed by some script with one or more sources and targets involved. etlog calls this unit a _transaction_ since it represents a complete iteration of ETL.

The below states that _the value in column 4 on line 5 in the users.csv found at /path/to/users.csv on host 148.29.12.100_ was _inserted_ _in the email column in the row with an id 38 in the users table in the socialapp database on the host 148.29.12.101 on port 5236_ on _August 8, 2013 at 5:43:03 AM_ by _parse-users.py at revision a32f87cb_.

```javascript
{
    "timestamp": "2013-08-13T05:43:03.32344",
    "action": "update",
    "script": {
        "uri": "https://github.com/cbmi/project/blob/master/parse-users.py",
        "version": "a32f87cb"
    },
    "source": {
        "type": "delimited",
        "delimiter": ",",
        "uri": "148.29.12.100/path/to/users.csv",
        "name": "users.csv",
        "line": 5,
        "column": 4
    },
    "target": {
        "type": "relational",
        "uri": "148.29.12.101:5236",
        "database": "socialapp",
        "table": "users",
        "row": { "id": 38 },
        "column": "email"
    }
}
```

- `timestamp` - Timestamp of when the transaction occurred
- `script` - An object representing the script used that performed the ETL and produced this transaction.
- `source` - An object or array of objects representing the sources of data being used in the target output.
- `target` - An object or array of objects representing the targets 
- `action` - The action that was performed on the target. This may not be applicable or available depending on what operation the script is performing. The value is generally target-specific; for example, `insert` and `delete` for row-based operations, `append` for file-based writing, `pop` for a Redis list, etc.

### Script

Properties about the script used for this transction.

- `uri` - URI which denotes where the data lives, where came from (streamed) or where can be accessed in the future. The latter value is preferred since it theoretically enables future access.
- `version` - The version of the script. This could be a timestamp, Git commit SHA, tag version, etc. This useful when an issue is found with the script and it can be assessed whether a transaction is affected.
- `language` - The primary programming language the script is written in. If not defined, this will be attempted to be inferred from filename specified in `uri`.
- `code` - The actual code as text or statement used during this transaction. This is most useful for scripts that perform in-place operations that never expose the data itself. For example, a SQL statement that selects data from one table and inserts it into a new table.

### Store

The base store type which all other types extend.

- `uri` - URI which denotes where the data lives, where came from (streamed) or where can be accessed in the future. The latter value is preferred since it theoretically enables future access. The completeness and value of this is store dependent.
- `type` - The store type as a string. This is used when messages are parsed and for downstream processing. If not defined, the type will attempted to be inferred by keys supplied.
- `value` - The value or array of values processed. For sources this would typically be the pre-transformed data while for targets this would be post-processed data. Supplying the value here is typically unnecessary since the value _could_ be accessed using the other information supplied in store data. However, for systems that treat this as an audit log or if the target system does not perform an versioning of data of it's own, this could act as primitive store of values.

### Relational

Extends `Store`

Store representing a data based on the relational model. Examples include PostgreSQL, SQLite, Oracle, MySQL.

- `database` - The database name
- `schema` - The schema name for databases that support them.
- `table` - The table name
- `row` - Object containing the lookup of the row
- `column` - The column name

### File

Extends `Store`

The base type for file-based stores.

- `name` - The name of the file. If the `uri` is supplied with a path, the name of the file will be extracted if not supplied.

### Binary

Extends `File`

Store for binary files or text files which are not line-based. The location of the data is defined by byte positions and ranges.

- `bytes` - The byte position, range or series of bytes and ranges.
    - `130`
    - `130-150`
    - `120-127,140-150`

### Text

Extends `File`

Store for text files. The location of the data is defined by lines and chars.

- `line` - The line, line range or series of lines and ranges.
    - `1`
    - `1-4`
    - `1-2,5,7-9`
- `chars` - The same format as `bytes` above except this is relative to the target line.

### Delimited

Extends `Text`

Store for delimited text files. The location is denoted by the `line` and `column` fields. Examples includes CSV and tab-delimited files.

- `delimiter` - The delimiter used which denotes the structure.
- `header` - A boolean, line number, or the header line of the file delimited by the above delimiter. Note, this should only be supplied if the columns are guaranteed consistent for all rows in the file.
- `column` - The column index or name (if a `header` exists) where the data is defined. For indexes, this can be a range. For column names this can be a comma-separated list of names.
    - `1,3-5,8`
    - `id,email,username`

### Spreadsheet

Extend `File`

Store for spreadsheet files such as Microsoft Excel, Google Docs Spreadsheet, and OpenOffice Spreadsheet. Data is tabular

- `sheet` - The index or name of the sheet
- `header` - A boolean, line number, or the header line of the file delimited by the above delimiter. Note, this should only be supplied if the columns are guaranteed consistent for all rows in the file.
- `row` - The line, line range or series of lines and ranges.
- `column` - The column index or name (if a `header` exists) where the data is defined. For indexes, this can be a range. For column names this can be a comma-separated list of names.
    - `1,3-5,8`
    - `id,email,username`

### Document

Extends `Store`

Store for JSON-based data

- `path` - A forward slash-delimited path to the value. This can be a single path or array of paths.
    - `person/addresses/0/city`

### MongoDB

Extends `Document`

Store for documents stored in MongoDB databases.

- `database` - The database name. If not supplied the database name will attempt to be extracted from the `uri`.
- `collection` - The collection name where the document is stored.

### Key/Value

Extends `Store`

For simple key/value-based stores, the keys that were used.

- `key` - The key or array of keys to the values.

## Use Cases

- Timeline view of data
    - What sources are used to populate a table/database?
    - How frequent do the records change?
- Audit trail
    - How and when the data was populated
    - Before and after values of transformation (if `value` is populated)
