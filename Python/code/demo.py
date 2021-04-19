# import sys
# import getopt

# try:
#     opts, args = getopt.getopt(sys.argv[1:], "ho:v", ["help", "output="])
# except getopt.GetoptError as err:
#     print(str(err)) # 此处输出类似"option -a not recognized"的出错信息
#     usage()
#     sys.exit(2)
# output = None
# verbose = False
# for o, a in opts:
#     if o == "-v":
#         verbose = True
#     elif o in ("-h", "--help"):
#         usage()
#         sys.exit()
#     elif o in ("-o", "--output"):
#         output = a
#     else:
#         assert False, "unhandled option"


# from optparse import OptionParser

# parser = OptionParser()
# parser.add_option("-f", "--file", dest="fileanme", help="write report to FILE", metavar="FILE")
# parser.add_option("-q", "--quiet", action="store_false", dest="verbose", default=True, help="don't print status messages to stdout")
# (options, args) = parser.parse_args()


# import argparse

# parser = argparse.ArgumentParser()
# parser.add_argument('-o', '--output')
# parser.add_argument('-v', dest='verbose', action='store_true')
# args = parser.parse_args()


from pandas import Series
obj1 = Series([1, 'a', (1, 2), 3], index=['a', 'b', 'c', 'd'])

Series({"Book": "Python", "Author": "Dan", "ISBN": "011334", "Price": 25}, index=['book', 'Author', 'ISBM', 'Price'])