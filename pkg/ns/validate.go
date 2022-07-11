package ns

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrNoNamespaceToRename   = errors.New("no namespace to rename")
	ErrNamespaceInvalid      = errors.New("invalid namespace")
	ErrCanHaveOnlyOneDefault = errors.New("can have only one default namespace")
	ErrNotYetSupported       = errors.New("not yet supported")

	ErrNoFileToRename = errors.New("no file to rename")
)

var rfc1123re = regexp.MustCompile(`^([[:alnum:]][[:alnum:]\-]{0,61}[[:alnum:]]|[[:alpha:]])$`)

func (r Rename) Validate() error {
	if len(r.Namespaces) == 0 {
		return ErrNoNamespaceToRename
	}
	if len(r.Inputs) == 0 {
		return ErrNoFileToRename
	}
	anyFound := false
	for _, namespace := range r.Namespaces {
		if err := namespace.Validate(); err != nil {
			return err
		}
		if namespace.From == AnyNamespace {
			if anyFound {
				return ErrCanHaveOnlyOneDefault
			}
			anyFound = true
		}
	}

	return r.validateUnsupported()
}

// TODO: remove unsupported validations, when the are implemented
func (r Rename) validateUnsupported() error {
	if len(r.Namespaces) > 1 {
		return fmt.Errorf("%w: more then one namespace", ErrNotYetSupported)
	}
	if r.Namespaces[0].From != AnyNamespace {
		return fmt.Errorf("%w: specific namespace to replace", ErrNotYetSupported)
	}
	return nil
}

func (r NamespaceRename) Validate() error {
	if !r.isValid() {
		return fmt.Errorf("%w: %s", ErrNamespaceInvalid, r)
	}
	return nil
}

func (r NamespaceRename) isValid() bool {
	return (r.From == AnyNamespace || rfc1123re.MatchString(r.From)) &&
		rfc1123re.MatchString(r.To)
}

func (r NamespaceRename) String() string {
	if r.From == AnyNamespace {
		return r.To
	}
	return fmt.Sprintf("%s=%s", r.From, r.To)
}
