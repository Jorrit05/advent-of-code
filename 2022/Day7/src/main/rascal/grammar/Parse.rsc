module grammar::Parse


import grammar::Syntax;
import ParseTree;

FileSystem parseFs(loc file) = parse(#start[FileSystem], file).top;
