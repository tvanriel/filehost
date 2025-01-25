# Filehost

Filehost is a simple S3-Backed File Hosting Service for an internet community.

Just place it behind an authentication proxy such as Authelia.  give it write access to a S3 Bucket and you're done.

everything is uploaded using Presign URLs.  So your S3 Bucket does the heavy lifting for you.
