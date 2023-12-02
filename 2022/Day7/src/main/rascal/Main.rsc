module Main

import IO;
import lib::Debug;
import grammar::Load;
import grammar::AST;


tuple[str,int] getFolderSize(str folderName, FileSystem ast) {
    int size = 0;
    bottom-up-break visit(ast){
        case dir: cd(str name, list[FolderContents] fs): {
            if(name == folderName) {
                // println("subfol  dername <name>");
                for (content <- fs) {
                    if (content.size?){
                        size += content.size;
                    }
                }
                for (content <- fs) {
                    if (content.size?){
                        continue;
                    } else {
                        // if (size > 100000) {
                        //     return <folderName, 0>;
                        // }
                        // it's a directory
                        <name,x> = getFolderSize(content.name, ast);
                        size += x;
                    }
                }
            }
        }
    }
    return <folderName,size>;
}

void main(){
    loc fileLoc = |project://Day7/input.txt|;
    ast = load(fileLoc);

    list[tuple[str,int]] folderSizes = [];
    list[str] names = [];

    bottom-up visit(ast) {
        case directory(str name): {
            names += name;
        }
    }

    // visit(ast){
    //     case cd(str name, list[FolderContents] fs): {
    //         // println("Original name: <name>");
    //         int fileSizes = 0;
    //         tuple[str, int] subFolders = <"",0>;
    //         for (content <- fs) {
    //             if (content.size?){
    //                 fileSizes += content.size;
    //             } else {
    //                 // it's a directory
    //                 subFolders = getFolderSize(content.name, ast);
    //             }
    //         }
    //         <_, size> = subFolders;
    //         // println("Original FileSizes: <fileSizes>");
    //         // println("   - SubFolder Sizes: <size>");
    //         // println("   TOTAL <name>:  <size+fileSizes>");

    //         folderSizes += <name, size+fileSizes>;


    //     }
    // }

    // println(folderSizes);
    int result = 0;
    for (<name, size> <- folderSizes) {
        if(size <= 100000) {
            println("<name>: <size>");
            result += size;
        }
    }
    println(names);
     dumpToJson("out.json", ast);
}

// void main() {
//     loc file = |project://Day7/test.txt|;
//     list[str] fileLines = readFileLines(file);
//     FileSystem fileSystem = Root("/");
//     FileSystem parent = fileSystem;

//     for (line <- fileLines) {
//         if(/\$ cd \// := line) {
//             // fileSystem = Directory("/");
//             continue;
//         }

//         if(/^\$\ cd <folder:.*$>/  := line) {
//             println(folder);
//             // newDir = addDir(folder, parent, []);

//         }
//         // else if (/\$ ls/ := line) {
//         //     println("ls");
//         // }
//     }
//     // x = directory("b", y, root("/"));
//     // y = directory("a", x);

//     // visit(fileSystem){
//     //     case xx: _ : println("<xx>--");
//     // }
//     dumpToJson("out.json", fileSystem);
// }
