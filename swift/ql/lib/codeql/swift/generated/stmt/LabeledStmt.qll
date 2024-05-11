// generated by codegen/codegen.py
/**
 * This module provides the generated definition of `LabeledStmt`.
 * INTERNAL: Do not import directly.
 */

private import codeql.swift.generated.Synth
private import codeql.swift.generated.Raw
import codeql.swift.elements.stmt.Stmt

/**
 * INTERNAL: This module contains the fully generated definition of `LabeledStmt` and should not
 * be referenced directly.
 */
module Generated {
  /**
   * INTERNAL: Do not reference the `Generated::LabeledStmt` class directly.
   * Use the subclass `LabeledStmt`, where the following predicates are available.
   */
  class LabeledStmt extends Synth::TLabeledStmt, Stmt {
    /**
     * Gets the label of this labeled statement, if it exists.
     */
    string getLabel() {
      result = Synth::convertLabeledStmtToRaw(this).(Raw::LabeledStmt).getLabel()
    }

    /**
     * Holds if `getLabel()` exists.
     */
    final predicate hasLabel() { exists(this.getLabel()) }
  }
}
