package httprequests

type GithubIssueBody struct {
	Owner     string `json:"owner"`
	Repo      string `json:"repo"`
	Body      string `json:"body"`
	IsBug     bool   `json:"isBug"`
	Email     string `json:"email"`
	BattleTag string `json:"battleTag"`
	Rating    string `json:"rating"`
}

func (b *GithubIssueBody) GetTitle() string {
	if b.IsBug {
		return "Bug report"
	} else {
		return "Feedback report"
	}
}

func (b *GithubIssueBody) GetLabels() []string {
	if b.IsBug {
		return []string{"bug"}
	} else {
		return []string{"feedback"}
	}
}

func (b *GithubIssueBody) GetBody() string {
	if b.IsBug {
		return b.getBugBody()
	} else {
		return b.getFeebackBody()
	}
}

func (b *GithubIssueBody) getBugBody() string {
	return `## Description

` + b.Body + `

---

## Attachments (Optional)

---

## Contact Information (Optional)

- **Email**: ` + b.Email + `
- **BattleTag**: ` + b.BattleTag
}

func (b *GithubIssueBody) getFeebackBody() string {
	return `## Description

` + b.Body + `

---

## Attachments (Optional)

---

## Contact Information (Optional)

- **Email**: ` + b.Email + `
- **BattleTag**: ` + b.BattleTag + `

---

## Rating

**Rating**: ` + b.Rating
}
