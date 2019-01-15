workflow attach {
  typespace => 'aws',
  input => (
    Hash[String,String] $tags = lookup('aws.tags'),
  ),
  output => (
    String $vpc_id,
  )
} {
  resource vpc {
    input  => ($tags),
    output => ($vpc_id)
  }{
    cidr_block => '192.168.0.0/16',
  }
}
