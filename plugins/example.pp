type Example::Address = {
  attributes => {
    line_one => { type=>String, value=>'' },
  }
}

type Example::Person = {
  attributes => {
    age => { type=>Integer, value=>0 },
    name => { type=>String, value=>'' },
    human => { type=>Boolean, value=>false },
    address => { type=>Optional[Example::Address], value=>undef}
  }
}

type Example::Three = {
  attributes => {
    name => { type=>String, value=>'' },
  }
}

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
  }
}

