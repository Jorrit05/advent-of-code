module grammar::Load

import grammar::Syntax;
import grammar::Parse;
import grammar::AST;
import lib::Debug;

import ParseTree;
import IO;

grammar::AST::FileSystem implode(grammar::Syntax::FileSystem fs) = implode(#FileSystem, fs);

grammar::AST::FileSystem load(loc l) = implode(parseFs(l));
