API Calls
=========

Add a File
----------
/siafile/ POST

Send JSON:
```
{
  Siafile: STRING,
  Title: STRING,
  Tags: [STRING]
}
```

Returns:
```
{
  Success: BOOL
}
```

Remove A File
-------------
/siafile/ DELETE
Send JSON:
Send a hash of the Sia file and credentials using basic auth. Note only Admins will have delete priviliges at first.
And the hash is of the Sia file.

```
{
  Hash: STRING,
  Password: STRING
}
```

Returns:

```
{
  Success: BOOL
}
```

Querying for a file
-------------------

Search

/siafile/search/ GET

Parameters has list of strings to query by


Returns:

```
{
  Hash: STRING,
  Title: STRING,
  Tag: [STRING]
}
```

Get Sia File
------------
/siafile/<HASH> GET

Returns:

```
{
  Hash: STRING,
  Title: STRING,
  Tag: [STRING],
  Siafile: STRING
}
```
