module grammar::Syntax

layout Whitespace = [\ \t\n\r]* !>> [\ \t\n\r];

lexical DirName = [/a-zA-Z0-9]+ !>> [/a-zA-Z0-9];
lexical ID = [.a-zA-Z0-9]+ !>> [.a-zA-Z0-9];

lexical Integer = [0-9]+ !>> [0-9]
            ;

start syntax FileSystem =
                filesystem: DrStructure dr
                ;

syntax DrStructure = Root r SubDirs subdirs*;

syntax Root =  "$ cd /";
syntax SubDirs
            = cd: "$ cd " DirName dirName "$ ls" FolderContents* folderContents
            | up: "$ cd .."
            ;

syntax FolderContents
        = file: Integer size ID fileName
        | directory: "dir" ID directoryName
        ;
