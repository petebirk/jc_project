#!/bin/bash
echo "** TESTING /hash POSTS **"
key1=`curl -s -X POST -d "password=angryMonkey1"  http://localhost:8080/api/v1/hash`
echo "key1 = $key1"
key2=`curl -s -X POST -d "password=angryMonkey2"  http://localhost:8080/api/v1/hash`
echo "key2 = $key2"
key3=`curl -s -X POST -d "password=angryMonkey3"  http://localhost:8080/api/v1/hash`
echo "key3 = $key3"
key4=`curl -s -X POST -d "password=angryMonkey4"  http://localhost:8080/api/v1/hash`
echo "key4 = $key4"
key5=`curl -s -X POST -d "password=angryMonkey5"  http://localhost:8080/api/v1/hash`
echo "key5 = $key5"

echo "** TESTING /hash/:requestId GETS before 5 sec **"
value1=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key1`
echo "value1 = $value1"
value2=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key2`
echo "value2 = $value2"
value3=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key3`
echo "value3 = $value3"
value4=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key4`
echo "value4 = $value4"
value5=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key5`
echo "value5 = $value5"

echo "Sleeping 5 seconds..."
sleep 5

echo "** TESTING /hash/:requestId GETS after 5 sec **"
value1=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key1`
echo "value1 = $value1"
value2=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key2`
echo "value2 = $value2"
value3=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key3`
echo "value3 = $value3"
value4=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key4`
echo "value4 = $value4"
value5=`curl -s -X GET  http://localhost:8080/api/v1/hash/$key5`
echo "value5 = $value5"

echo "** TESTING /hash/:requestId GET with invalid negative requestId value **"
returnCode=`curl -s -X GET  http://localhost:8080/api/v1/hash/-5`
echo "Returned $returnCode"

echo "** TESTING /hash POST with invalid password string **"
returnCode=`curl -s -X POST -d "password=<script>alert!</script>"  http://localhost:8080/api/v1/hash`
echo "Returned $returnCode"


echo "** TESTING /stats valid GET return. **"
returnCode=`curl -s -X GET  http://localhost:8080/api/v1/stats`
echo "Returned $returnCode"

echo "** TESTING /stats invalid POST. **"
returnCode=`curl -s -X GET  http://localhost:8080/api/v1/stats`
echo "Returned $returnCode"

echo "** TESTING bad password property -d (passwo=angryMonkey1) in POST **"
returnCode=`curl --write-out %{http_code} --silent --output /dev/null -d "passwo=angryMonkey1"  http://localhost:8080/api/v1/hash`
echo "Returned $returnCode"

echo "** TESTING 202 code (Accepted) for valid POST **"
returnCode=`curl --write-out %{http_code} --silent --output /dev/null -d "password=angryMonkey1"  http://localhost:8080/api/v1/hash`
echo "Returned $returnCode"

echo "** TESTING 404 code (NotFound) for invalid URL /api/v1/ **"
returnCode=`curl --write-out %{http_code} --silent --output /dev/null -d "password=angryMonkey1"  http://localhost:8080/api/v1/`
echo "Returned $returnCode"

echo "** TESTING 405 code (MethodNotAllowed) for invalid Method PUT to /hash **"
returnCode=`curl --write-out %{http_code} --silent --output /dev/null -X PUT -d "password=angryMonkey1"  http://localhost:8080/api/v1/hash`
echo "Returned $returnCode"

echo "** TESTING 405 code (MethodNotAllowed) for invalid Method GET to /hash **"
returnCode=`curl --write-out %{http_code} --silent --output /dev/null -X GET http://localhost:8080/api/v1/hash`
echo "Returned $returnCode"

echo "** ADDING /hash POSTS before shutdown **"
key1=`curl -s -X POST -d "password=angryMonkey1"  http://localhost:8080/api/v1/hash`
echo "key1 = $key1"
key2=`curl -s -X POST -d "password=angryMonkey2"  http://localhost:8080/api/v1/hash`
echo "key2 = $key2"
key3=`curl -s -X POST -d "password=angryMonkey3"  http://localhost:8080/api/v1/hash`
echo "key3 = $key3"
key4=`curl -s -X POST -d "password=angryMonkey4"  http://localhost:8080/api/v1/hash`
echo "key4 = $key4"
key5=`curl -s -X POST -d "password=angryMonkey5"  http://localhost:8080/api/v1/hash`
echo "key5 = $key5"
echo "** TESTING Shutdown. **"
returnCode=`curl -s -X POST http://localhost:8080/api/v1/shutdown`
