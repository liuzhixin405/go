using Hprose.Client;

namespace go_rpc
{
    internal class Program
    {
        static void Main(String[] args)
        {
            TestMethod();
            Console.ReadKey();
        }

        static void TestMethod()
        {
            HproseHttpClient client = new HproseHttpClient(" http://192.168.253.130:8080/");
            client.KeepAlive = true;
            Console.WriteLine(client.Invoke("hello, csharp client ", new Object[] { "test hello" }));
            Console.ReadLine();
        }
    }
}