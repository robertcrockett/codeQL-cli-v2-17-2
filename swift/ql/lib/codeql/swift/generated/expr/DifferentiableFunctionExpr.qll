// generated by codegen/codegen.py
/**
 * This module provides the generated definition of `DifferentiableFunctionExpr`.
 * INTERNAL: Do not import directly.
 */

private import codeql.swift.generated.Synth
private import codeql.swift.generated.Raw
import codeql.swift.elements.expr.ImplicitConversionExpr

/**
 * INTERNAL: This module contains the fully generated definition of `DifferentiableFunctionExpr` and should not
 * be referenced directly.
 */
module Generated {
  /**
   * INTERNAL: Do not reference the `Generated::DifferentiableFunctionExpr` class directly.
   * Use the subclass `DifferentiableFunctionExpr`, where the following predicates are available.
   */
  class DifferentiableFunctionExpr extends Synth::TDifferentiableFunctionExpr,
    ImplicitConversionExpr
  {
    override string getAPrimaryQlClass() { result = "DifferentiableFunctionExpr" }
  }
}
