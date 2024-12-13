// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:camt.060.001.05
package AccountReportingRequest_camt_060_001_05

import (
	"github.com/moov-io/base"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Element validations

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	fedwire.AddError(&errs, baseName+".AcctRptgReq", v.AcctRptgReq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations

func (v AccountIdentification4Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountIdentification4Choice1"
	if v.Othr != nil {
		fedwire.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AccountReportingRequestV05) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountReportingRequestV05"
	fedwire.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	fedwire.AddError(&errs, baseName+".RptgReq", v.RptgReq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification61) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification61"
	fedwire.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccount381) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccount381"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	fedwire.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccountType2Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccountType2Choice1"
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemIdentification2Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemIdentification2Choice1"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification21) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification21"
	fedwire.AddError(&errs, baseName+".ClrSysId", v.ClrSysId.Validate())
	fedwire.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification181) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification181"
	fedwire.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if v.Othr != nil {
		fedwire.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericAccountIdentification11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericAccountIdentification11"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericFinancialIdentification11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericFinancialIdentification11"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GroupHeader771) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GroupHeader771"
	fedwire.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	fedwire.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40Choice1"
	if v.Agt != nil {
		fedwire.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ReportingRequest51) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ReportingRequest51"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	fedwire.AddError(&errs, baseName+".ReqdMsgNmId", v.ReqdMsgNmId.Validate())
	if v.Acct != nil {
		fedwire.AddError(&errs, baseName+".Acct", v.Acct.Validate())
	}
	fedwire.AddError(&errs, baseName+".AcctOwnr", v.AcctOwnr.Validate())
	if v.RptgSeq != nil {
		fedwire.AddError(&errs, baseName+".RptgSeq", v.RptgSeq.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SequenceRange1Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SequenceRange1Choice1"
	if v.FrToSeq != nil {
		fedwire.AddError(&errs, baseName+".FrToSeq", v.FrToSeq.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SequenceRange11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SequenceRange11"
	fedwire.AddError(&errs, baseName+".FrSeq", v.FrSeq.Validate())
	fedwire.AddError(&errs, baseName+".ToSeq", v.ToSeq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v AccountReportingFedwireFunds1) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "DTLS", "DTLR", "ETOT", "ABAR"); err != nil {
		return err
	}
	return nil
}

func (v AccountTypeFRS1) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "S", "M"); err != nil {
		return err
	}
	return nil
}

func (v EndpointIdentifierFedwireFunds1) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[A-Z0-9]{8,8}`); err != nil {
		return err
	}
	if err := fedwire.ValidateMinLength(string(v), 8); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 8); err != nil {
		return err
	}
	return nil
}

func (v ExternalClearingSystemIdentification1CodeFixed) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "USABA"); err != nil {
		return err
	}
	return nil
}

func (v Max35Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v MessageNameIdentificationFRS1) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`); err != nil {
		return err
	}
	if err := fedwire.ValidateLength(string(v), 15); err != nil {
		return err
	}
	return nil
}

func (v RoutingNumberFRS1) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[0-9]{9,9}`); err != nil {
		return err
	}
	if err := fedwire.ValidateLength(string(v), 9); err != nil {
		return err
	}
	return nil
}

func (v XSequenceNumberFedwireFunds1) Validate() error {
	if err := fedwire.ValidateMinInclusive(float64(v), 000001); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxInclusive(float64(v), 999999); err != nil {
		return err
	}
	return nil
}
