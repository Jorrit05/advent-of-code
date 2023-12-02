module Day8 where

import Data.List

northSouth :: [String] -> [String]
northSouth = transpose

eastWest :: [String] -> [String]
eastWest = map reverse

southNorth :: [String] -> [String]
southNorth = transpose . reverse

checkTrees :: [String] -> Int -> [Bool]
checkTrees treeLine = undefined

orList :: [Bool] -> [Bool] -> [Bool]
orList [] [] = []
orList _ [] = []
orList [] _ = []
orList (x:xs) (y:ys) = (x || y): orList xs ys

higherThen :: Char -> [Char] -> Bool
higherThen x = any (> x)

mytest :: String -> [Bool]
mytest str = map (\x -> higherThen x str) str

nextIsGreater xs = [x | (x:y:_) <- tails xs, x<y]

higherThenAllPrevious :: Char -> String  -> Bool
higherThenAllPrevious = undefined
gen f = xs where xs = map f $ inits xs
reachable :: String -> [Bool]
reachable str = map (`higherThenAllPrevious` str) str
            where
               stringSize = length str

main = do
   fileContent <- readFile "test.txt"
   let fileList = lines fileContent
   let r = map reachable fileList
   print $ southNorth fileList