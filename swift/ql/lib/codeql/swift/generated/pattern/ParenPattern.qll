// generated by codegen/codegen.py
/**
 * This module provides the generated definition of `ParenPattern`.
 * INTERNAL: Do not import directly.
 */

private import codeql.swift.generated.Synth
private import codeql.swift.generated.Raw
import codeql.swift.elements.pattern.Pattern

/**
 * INTERNAL: This module contains the fully generated definition of `ParenPattern` and should not
 * be referenced directly.
 */
module Generated {
  /**
   * INTERNAL: Do not reference the `Generated::ParenPattern` class directly.
   * Use the subclass `ParenPattern`, where the following predicates are available.
   */
  class ParenPattern extends Synth::TParenPattern, Pattern {
    override string getAPrimaryQlClass() { result = "ParenPattern" }

    /**
     * Gets the sub pattern of this paren pattern.
     *
     * This includes nodes from the "hidden" AST. It can be overridden in subclasses to change the
     * behavior of both the `Immediate` and non-`Immediate` versions.
     */
    Pattern getImmediateSubPattern() {
      result =
        Synth::convertPatternFromRaw(Synth::convertParenPatternToRaw(this)
              .(Raw::ParenPattern)
              .getSubPattern())
    }

    /**
     * Gets the sub pattern of this paren pattern.
     */
    final Pattern getSubPattern() {
      exists(Pattern immediate |
        immediate = this.getImmediateSubPattern() and
        if exists(this.getResolveStep()) then result = immediate else result = immediate.resolve()
      )
    }
  }
}