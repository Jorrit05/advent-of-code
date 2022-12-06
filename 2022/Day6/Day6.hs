module Day6 where

import Data.List ( nub )

duplicates :: String -> Bool
duplicates str = length (nub str) == length str

parseString :: String -> Int -> Int -> Int
parseString str counter method | duplicates $ take method str = counter + method
                               | otherwise = parseString (drop 1 str) (counter + 1) method
main = do
    fileContent <- readFile "input.txt"
    let dataStream =  head $ words fileContent
    print (parseString dataStream 0 4, parseString dataStream 0 14)