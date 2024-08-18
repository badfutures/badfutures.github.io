using System.Text.RegularExpressions;

Regex r = Example.Demo();

static partial class Example
{
  [GeneratedRegex("a*[ab]")]
  public static partial Regex Demo();
}
