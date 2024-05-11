// generated by codegen/codegen.py
import codeql.swift.elements
import TestUtils

from ConsumeExpr x, string hasType, Expr getSubExpr
where
  toBeTested(x) and
  not x.isUnknown() and
  (if x.hasType() then hasType = "yes" else hasType = "no") and
  getSubExpr = x.getSubExpr()
select x, "hasType:", hasType, "getSubExpr:", getSubExpr