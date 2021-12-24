# https://stackoverflow.com/questions/36138250/how-do-i-test-methods-using-boto3-with-moto
import boto
import unittest
from boto.s3.key import Key
from moto import mock_s3
import boto3

class TestS3Actor(unittest.TestCase):
    mock_s3 = mock_s3()

    def setUp(eslf):
        self.mock_s3.start()
        # Set to your region where your s3 bucket is on
        self.location = "ca-central-1"
        self.bucket_name = "firefile"
        self.key_name = "stats_com/fake_fake/test.json"
        self.key_contents = "Dummy test data"
        s3 = boto.connect_s3()
        bucket = s3.create_bucket(self.bucket_name, location=self.location)
        k = Key(bucket)
        k.key = self.key_name
        k.set_contents_from_string(self.key_contents)

    def tearDown(self):
        self.mock_s3.stop()

    def test_s3_boto3(self):
        s3 = boto3.resource('s3', region_name=self.location)
        bucket = s3.Bucket(self.bucket_name)
        assert bucket.name == self.bucket_name
        keys = list(bucket.objects.filter(Prefix=self.key_name))
        assert len(keys) == 1
        assert keys[0].key == self.key_name
        
        s3.Object(self.bucket_name, self.key_name).put(Body="new")
        key = s3.Object(self.bucket_name, self.key_name).get()
        assert 'new' == key["Body"].read()
