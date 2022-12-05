module Day5 where

import System.IO
import Control.Monad
import Data.List
import Data.Char

parseData :: String -> Int -> [Char]
parseData str counter | counter >= length str = []
                      | otherwise = str !! counter : parseData str (counter + 4)

readInt :: String -> Int
readInt x | null x = 0
          | otherwise = read x


parseMoves :: [String] -> [(Int, Int, Int)]
parseMoves [] = []
parseMoves (x:xs) = (head digits,  (digits !! 1) -1,  (digits !! 2) -1) : parseMoves xs
                     where
                        digits = map readInt $ filter (isDigit . head ) $ split x

replaceAtIndex :: Int -> a -> [a] -> [a]
replaceAtIndex n item xs = a ++ (item:b)
                            where
                                (a, _:b) = splitAt n xs

split :: String -> [String]
split str = case break (==' ') str of
                (a, ' ':b) -> a : split b
                (a, "")    -> [a]

move :: (Int, Int, Int) -> [String] -> [String]
move  (amount, from, to) stack = new
                            where
                              takenLetters = reverse $ take amount $ stack !! from
                              insert = replaceAtIndex to (takenLetters ++  stack !! to) stack
                              new = replaceAtIndex from (drop amount $ stack !! from) insert

-- Copy from move, only removed the 'reverse'
movePuz2 :: (Int, Int, Int) -> [String] -> [String]
movePuz2  (amount, from, to) stack = new
                            where
                              takenLetters =  take amount $ stack !! from
                              insert = replaceAtIndex to (takenLetters ++  stack !! to) stack
                              new = replaceAtIndex from (drop amount $ stack !! from) insert

main = do
    fileContent <- readFile "input.txt"
    let fileList = lines fileContent
    let filteredStack = map (`parseData` 1) $ takeWhile (notElem '1' ) fileList
    let startStack  = map (dropWhile isSpace) $ transpose filteredStack
    let moves = parseMoves $ dropWhile (notElem 'm' ) fileList

    let result1 = foldl (flip move) startStack moves
    let result2 = foldl (flip movePuz2) startStack moves

    print $ map head $ filter (/= "") result1 -- MQSHJMWNH
    print $ map head $ filter (/= "") result2 -- LLWJRBHVZ

