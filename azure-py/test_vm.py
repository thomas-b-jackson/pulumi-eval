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
    # Test if the service has tags and a name tag.
    @pulumi.runtime.test
    def test_server_tags(self):
        def check_tags(args):
            linux_virtual_machine = args
            self.assertIsEqual(linux_virtual_machine.location,'westus2')

        print(infra.server)
        return pulumi.Output.all(infra.linux_virtual_machine).apply(check_tags)

