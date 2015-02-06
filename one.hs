import Text.ParserCombinators.Parsec
import System.IO

main :: IO ()
main = do
       logHandle <- openFile "access.log" ReadMode
       logLines <- hGetContents logHandle

       case parse logFile " " logLines of
         Left _ -> putStrLn "Error parsing input"
         Right r -> mapM_ (print . head) r
       hClose logHandle

logFile = endBy line eol

line :: GenParser Char st [String]
line = sepBy cell (char ' ')

cell :: GenParser Char st String
cell = many (noneOf " \n")

eol :: GenParser Char st Char
eol = char '\n'

parseLogLine :: String -> Either ParseError [[String]]
parseLogLine = parse logFile "(unknown)"
