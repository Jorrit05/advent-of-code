module Day9 where

import Data.Char
import Data.List


readInt :: String -> Int
readInt = read

split :: String -> [String]
split str = case break (==' ') str of
                (a, ' ':b) -> a : split b
                (a, "")    -> [a]

move :: Char -> Int -> (Int, Int, Int) -> [(Int, Int, Int)]
move   _ 0 currentLoc = []
move 'R' n (r, c, v) = (r + 1, c, v) : move 'R' (n-1) (r + 1, c, v)
move 'U' n (r, c, v) = (r, c + 1, v) : move 'U' (n-1) (r, c + 1, v)
move 'L' n (r, c, v) = (r - 1, c, v) : move 'L' (n-1) (r - 1, c, v)
move 'D' n (r, c, v) = (r, c - 1, v) : move 'D' (n-1) (r, c - 1, v)
move _ _ _ = []

handleMoves :: (Int, Int, Int) -> [(Char, Int)] -> [(Int, Int, Int)]
handleMoves _ []  = []
handleMoves currentLoc ((m, n):xs) = moves ++ handleMoves (last moves) xs
                                where
                                    moves = move m n currentLoc

followHead :: (Int, Int, Int) -> [(Int, Int, Int)]-> [(Int, Int, Int)]
followHead _ []  = []
followHead _ [x]  = []
followHead (a,b,c) ((m, n, v):(x, y, z):xs)  | x - a > 1 || y - b > 1 || a -x > 1 || b - y> 1 =  (a,b,c) : followHead (m, n, 1) ((x, y, z):xs)
                                             | otherwise  = followHead (a,b,c) ((x, y, z):xs)

puzzle2 :: (Int, Int, Int) -> [(Int, Int, Int)]-> [(Int, Int, Int)]
puzzle2 _ []  = []
puzzle2 (a,b,c) moveList | a == x && b == y = (a,b,1) :  puzzle2 (a,b,1) (drop 9 moveList)
                         | x - a > 1 || y - b > 1 || a -x > 1 || b - y> 1 =  (a,b,1) :  puzzle2 (x,y,1) (drop 9 moveList)
                         | otherwise =  (a,b,1) : puzzle2 (x,y,1) (drop 9 moveList)
                           where
                              newList = take 9 moveList
                              positions = tail $ followHead (a,b,c) newList
                              (x,y,z) = head positions


getMoves :: [String] -> [(Char, Int)]
getMoves = map (\ x -> (head x, readInt (last (split x))))

main = do
   fileContent <- readFile "test.txt"
   let moveList = handleMoves (0,0,0) $ getMoves $ lines fileContent

   let tailList = nub $ (0,0,1) : followHead (0,0,0) moveList
   let puzzleTwo = nub $ (0,0,1) : puzzle2 (0,0,0) moveList
   -- let puzzle2 = nub $ (0,0,1) : puzzle2 (0,0,0) moveList
   -- puzzle2 (0,0,0) moveList
   print $ length  puzzleTwo
   -- print $