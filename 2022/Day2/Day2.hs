module Day2 where

import System.IO
import Control.Monad
import Data.List

scoreMapping :: Char -> Int
scoreMapping 'A' = 1 -- Rock
scoreMapping 'B' = 2 -- Paper
scoreMapping 'C' = 3 -- Scissor
scoreMapping 'X' = 1 -- Rock
scoreMapping 'Y' = 2 -- Paper
scoreMapping 'Z' = 3 -- Scissor
scoreMapping _ = 0

getWinner :: Char -> Char -> Int
getWinner 'A' 'X' = 3
getWinner 'A' 'Y' = 6
getWinner 'B' 'Y' = 3
getWinner 'B' 'Z' = 6
getWinner 'C' 'Z' = 3
getWinner 'C' 'X' = 6
getWinner _ _ = 0

getMyMove :: Char -> Char -> Int
getMyMove 'A' 'X' = 3 + 0 -- X = lose = Scissor = 3
getMyMove 'A' 'Y' = 1 + 3 -- Y = draw = rock = 1
getMyMove 'A' 'Z' = 2 + 6 -- Z = win  = Paper = 2
getMyMove 'B' 'X' = 1 + 0
getMyMove 'B' 'Y' = 2 + 3
getMyMove 'B' 'Z' = 3 + 6
getMyMove 'C' 'X' = 2 + 0
getMyMove 'C' 'Y' = 3 + 3
getMyMove 'C' 'Z' = 1 + 6
getMyMove _ _ = 0

    -- X = lose = 0
    -- y = draw = 3
    -- z - win = 6

calcScores :: (Char, Char) -> Int
calcScores (x,y) = scoreMapping y + getWinner x y

calcScoresPuzzle2 :: (Char, Char) -> Int
calcScoresPuzzle2 (x,y) = getMyMove x y

getAllMoves :: [String] -> [(Char, Char)]
getAllMoves moveList = [(head move, last move) | move <- moveList]

main = do
    fileContent <- readFile "input.txt"
    let moveList = getAllMoves $ lines fileContent
    let puzzle1 = sum $ map calcScores moveList
    let puzzle2 = sum $ map calcScoresPuzzle2 moveList

    return (puzzle1, puzzle2)
