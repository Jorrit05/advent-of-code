module Day1 where

import System.IO
import Control.Monad


readInt :: String -> Int
readInt x | null x = 0
          | otherwise = read x

parseFoodCarriers :: [String] -> [[Int]]
parseFoodCarriers [] = []
parseFoodCarriers foodCarriers =  slice : parseFoodCarriers (drop (length slice + 1) foodCarriers)
                                where
                                    slice = map readInt $ takeWhile ((> 0) . readInt) foodCarriers

qsort :: Ord a => [a] -> [a]
qsort [] = []
qsort (x:xs) = qsort smaller ++ [x] ++ qsort larger
               where
                larger = [a | a <- xs, a >= x]
                smaller = [a | a <- xs, a < x]

main = do
    fileContent <- readFile "input.txt"
    let carrierList = parseFoodCarriers $ lines fileContent
    let sumList = map sum carrierList
    return (maximum sumList, sum $ take 3 $ reverse $ qsort sumList)
