// generated by codegen/codegen.py
/**
 * This module provides the generated definition of `ExtensionDecl`.
 * INTERNAL: Do not import directly.
 */

private import codeql.swift.generated.Synth
private import codeql.swift.generated.Raw
import codeql.swift.elements.decl.Decl
import codeql.swift.elements.decl.GenericContext
import codeql.swift.elements.decl.NominalTypeDecl
import codeql.swift.elements.decl.ProtocolDecl

/**
 * INTERNAL: This module contains the fully generated definition of `ExtensionDecl` and should not
 * be referenced directly.
 */
module Generated {
  /**
   * INTERNAL: Do not reference the `Generated::ExtensionDecl` class directly.
   * Use the subclass `ExtensionDecl`, where the following predicates are available.
   */
  class ExtensionDecl extends Synth::TExtensionDecl, GenericContext, Decl {
    override string getAPrimaryQlClass() { result = "ExtensionDecl" }

    /**
     * Gets the extended type declaration of this extension declaration.
     */
    NominalTypeDecl getExtendedTypeDecl() {
      result =
        Synth::convertNominalTypeDeclFromRaw(Synth::convertExtensionDeclToRaw(this)
              .(Raw::ExtensionDecl)
              .getExtendedTypeDecl())
    }

    /**
     * Gets the `index`th protocol of this extension declaration (0-based).
     */
    ProtocolDecl getProtocol(int index) {
      result =
        Synth::convertProtocolDeclFromRaw(Synth::convertExtensionDeclToRaw(this)
              .(Raw::ExtensionDecl)
              .getProtocol(index))
    }

    /**
     * Gets any of the protocols of this extension declaration.
     */
    final ProtocolDecl getAProtocol() { result = this.getProtocol(_) }

    /**
     * Gets the number of protocols of this extension declaration.
     */
    final int getNumberOfProtocols() { result = count(int i | exists(this.getProtocol(i))) }
  }
}