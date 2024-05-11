// generated by codegen/codegen.py
/**
 * This module provides the generated definition of `PlatformVersionAvailabilitySpec`.
 * INTERNAL: Do not import directly.
 */

private import codeql.swift.generated.Synth
private import codeql.swift.generated.Raw
import codeql.swift.elements.AvailabilitySpec

/**
 * INTERNAL: This module contains the fully generated definition of `PlatformVersionAvailabilitySpec` and should not
 * be referenced directly.
 */
module Generated {
  /**
   * An availability spec based on platform and version, for example `macOS 12` or `watchOS 14`
   * INTERNAL: Do not reference the `Generated::PlatformVersionAvailabilitySpec` class directly.
   * Use the subclass `PlatformVersionAvailabilitySpec`, where the following predicates are available.
   */
  class PlatformVersionAvailabilitySpec extends Synth::TPlatformVersionAvailabilitySpec,
    AvailabilitySpec
  {
    override string getAPrimaryQlClass() { result = "PlatformVersionAvailabilitySpec" }

    /**
     * Gets the platform of this platform version availability spec.
     */
    string getPlatform() {
      result =
        Synth::convertPlatformVersionAvailabilitySpecToRaw(this)
            .(Raw::PlatformVersionAvailabilitySpec)
            .getPlatform()
    }

    /**
     * Gets the version of this platform version availability spec.
     */
    string getVersion() {
      result =
        Synth::convertPlatformVersionAvailabilitySpecToRaw(this)
            .(Raw::PlatformVersionAvailabilitySpec)
            .getVersion()
    }
  }
}