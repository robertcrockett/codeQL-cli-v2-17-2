// generated by codegen/codegen.py
/**
 * This module provides the generated definition of `PrefixUnaryExpr`.
 * INTERNAL: Do not import directly.
 */

private import codeql.swift.generated.Synth
private import codeql.swift.generated.Raw
import codeql.swift.elements.expr.ApplyExpr

/**
 * INTERNAL: This module contains the fully generated definition of `PrefixUnaryExpr` and should not
 * be referenced directly.
 */
module Generated {
  /**
   * INTERNAL: Do not reference the `Generated::PrefixUnaryExpr` class directly.
   * Use the subclass `PrefixUnaryExpr`, where the following predicates are available.
   */
  class PrefixUnaryExpr extends Synth::TPrefixUnaryExpr, ApplyExpr {
    override string getAPrimaryQlClass() { result = "PrefixUnaryExpr" }
  }
}