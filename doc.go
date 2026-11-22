// Package strreq validates strings against declarative length and
// character-class requirements. A Requirements value describes the
// constraints — minimum/maximum length, letter/lower/upper/number
// minima, and allowed or forbidden characters — and IsValid reports
// whether a string satisfies all of them. Typical consumers load the
// requirements from configuration and check user-facing fields such
// as usernames and passwords.
package strreq
