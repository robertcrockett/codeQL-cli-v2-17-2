import python

class ImplicitConcat extends StringLiteral {
  ImplicitConcat() { exists(this.getAnImplicitlyConcatenatedPart()) }
}

from StringLiteral s, boolean isConcat
where
  s instanceof ImplicitConcat and isConcat = true
  or
  not s instanceof ImplicitConcat and isConcat = false
select s.getLocation().getStartLine(), s.getText(), isConcat, s.getText().length(),
  s.getLocation().getStartColumn(), s.getLocation().getEndColumn()
