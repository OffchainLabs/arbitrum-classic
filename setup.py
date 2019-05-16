from setuptools import setup

setup(name='arbc-solidity',
      version='0.1',
      description='Compiler from solidity to AVM bytecode',
      url='http://github.com/OffchainLabs/arbc-solidity',
      author='Harry Kalodner',
      author_email='harry@offchainlabs.com',
      license='UNLICENSED',
      packages=['arbitrum', 'arbitrum.std', 'arbitrum.evm'],
      zip_safe=False,
      scripts=['bin/arbc-truffle-compile', 'bin/arbc-compile'],
      test_suite='nose.collector',
      tests_require=['nose']
)
