module Day2 where

import System.IO
import Control.Monad
import Data.List

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

data Game = Rock Int
            | Paper Int
            | Scissor Int
            deriving Show


instance Eq Game where
    Rock _ == Rock _= True
    Scissor _ == Scissor _ = True
    Paper _ == Paper _ = True
    _ == _ = False

instance Ord Game where
    (>) (Rock _) (Scissor _) = True
    (>) (Scissor _) (Paper _) = True
    (>) (Paper _)   (Rock _) = True
    (>) _ _ = False

get :: Game -> Int
get (Rock x) = x
get (Paper x) = x
get (Scissor x) = x

move :: Char -> Game
move x  | x == 'A' || x == 'X' = Rock 1
        | x == 'B' || x == 'Y' = Paper 2
        | otherwise = Scissor 3

    -- X = lose = 0
    -- y = draw = 3
    -- z - win = 6

getWinner :: Game -> Game -> Int
getWinner x y | x > y = 0
              | x == y = 3
              | otherwise = 6

calcScores :: (Char, Char) -> Int
calcScores (x,y) = get (move y) + getWinner (move x) (move y)

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
    -- (9241,14610)