import unittest
import pulumi

class MyMocks(pulumi.runtime.Mocks):
    def new_resource(self, type_, name, inputs, provider, id_):
        return [name + '_id', inputs]
    def call(self, token, args, provider):
        return {}

pulumi.runtime.set_mocks(MyMocks())

# Now actually import the code that creates resources, and then test it.
import infra

class TestingWithMocks(unittest.TestCase):
    @pulumi.runtime.test
    def test_subnet_id(self):
        def check_id(args):
            subnet = args
            self.assertIsNotNull(subnet.id)

        print(infra.server)
        return pulumi.Output.all(infra.subnet).apply(check_id)

