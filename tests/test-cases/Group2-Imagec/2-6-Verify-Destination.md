Test 2-6 - Verify Destination
=======

#Purpose:
To verify that when imagec is run with the -destination flag, then it saves the image where specified

#References:
* imagec --help

#Environment:
Standalone test requires nothing but imagec to be built

#Test Steps:
1. Issue the following command:
* imagec -standalone -reference photon -destination foo

#Expected Outcome:
* Command should return success
* All the checksums for each image layer in the foo directory should match the manifest file

#Possible Problems:
Make sure that you run imagec on the same hard drive partition as /tmp, otherwise you will receive a cross device link error.