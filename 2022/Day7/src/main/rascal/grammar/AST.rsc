module grammar::AST

import grammar::Syntax;


data FileSystem = filesystem(list[Statements] fs)
                ;

data Statements
            = cd(str Folder, list[FolderContents] folderContents)
            // | ls(list[FolderContents] folderContents)
            | up()
            ;

data FolderContents
        = file(int size, str name)
        | directory(str name)
        ;