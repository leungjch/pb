from setuptools import setup, find_packages

setup(
    name='pb_leungjch',
    url='https://github.com/leungjch/pb',
    author='Justin Leung',
    packages=find_packages(),
    install_requires=['grpcio-tools'],
    version='0.1',
    license='MIT'
)