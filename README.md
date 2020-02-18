# masspf

## Goal
This is a fairly simple code, it gets domains from the stdin and try to see if there is an mx record set. If there is it will try to get a SPF record for the domain. If that is missing it will report it. If SPF is set it will print it.

## Usage
Usage of masspf:
  -p int
    	Size of the workers pool. (default 20)
  -snm
    	Prints domains that have no MX record set
  -p int
    	Size of the workers pool. (default 20)
  -snm
    	Prints domains that have no MX record set
