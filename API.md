API Calls
=========

Add a File
----------
/siafile/ POST

Send JSON:
```
{
  title: STRING,
  filename: STRING,
  description: STRING,
  content: STRING,
  tags: [STRING],
}
```

Returns:
```
{
  _id: bson.ObjectId,
  title: STRING,
  filename: STRING,
  description: STRING,
  content: STRING,
  tags: [STRING],
  uploadedTime: TIME,
}
```

Remove A File
-------------
/siafile/ DELETE
Send JSON:
Send a hash of the Sia file and credentials using basic auth. Note only Admins
will have delete priviliges at first.
And the hash is of the Sia file.

```
{
  id: STRING,
}
```

Returns:

```
{
  success: BOOL
}
```

Querying for a file
-------------------

Search

/siafile/search/ GET

Parameters query string to search by


Returns:

```
[
{
  _id: bson.ObjectId,
  title: STRING,
  filename: STRING,
  description: STRING,
  tags: [STRING],
  uploadedTime: TIME,
}
]
```

Get Sia File
------------
/siafile/<ID> GET

Returns:

```
{
  _id: bson.ObjectId,
  title: STRING,
  filename: STRING,
  description: STRING,
  content: STRING,
  tags: [STRING],
  uploadedTime: TIME,
}
```
