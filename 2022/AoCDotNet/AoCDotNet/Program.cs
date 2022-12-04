namespace AoCDotNet;

class Program
{
    static void Main(string[] args)
    {
        string filename = @"test.txt";
        var lines = File.ReadLines(filename);

        foreach (var line in lines)
        {
            Console.WriteLine(line);
        }

        Console.Write("Press ENTER to exit");
        Console.ReadLine();
    }
}
