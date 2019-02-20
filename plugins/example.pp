workflow sample {
  typespace => 'example',
  output => (
    String $name
  )
} {
  resource person {
    output => ($name)
  }{
    age => 28,
    name => 'Bob',
    human => false,
    addresses => [Example::Address(
      lineOne => '30 East 60th Street',
      # annotations => 
    )]
  }
}

workflow sample2 {
  typespace => 'example',
  input => (
    String $foo = lookup('foo', undef, undef, "foo"),
    String $bar = lookup('bar', undef, undef, "bar"),
    String $baz = lookup('baz', undef, undef, "baz")
  ),
  output => (
    String $foo,
    String $bar,
    String $baz
  )
} {
  resource person {
    output => ($name)
  }{
    age => 28,
    name => 'Bob',
    human => false,
  }
}
