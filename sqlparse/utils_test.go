package sqlparse

import "testing"

func TestParseSQL(t *testing.T) {
	var tests = []struct {
		input string
		want  alterQuery
	}{
		{"ALTER TABLE `GiftBox` ADD COLUMN `isPublished` INT NOT NULL DEFAULT 0;", alterQuery{Action: "ADD COLUMN `isPublished` INT NOT NULL DEFAULT 0", Table: "GiftBox"}},
		{"ALTER TABLE User DROP COLUMN totalRevenueEarned;", alterQuery{Action: "DROP COLUMN totalRevenueEarned", Table: "User"}},
		{`
			ALTER TABLE User DROP COLUMN totalRevenueEarned;
			`, alterQuery{Action: "DROP COLUMN totalRevenueEarned", Table: "User"}},
		{`
				UPDATE SubscriptionLog SET userVendorToken = "" WHERE userVendorToken IS NULL;
				`, alterQuery{}},
	}
	for _, test := range tests {
		if _, got := ParseAlterQuery(test.input); got != test.want {
			t.Error(got)
		}
	}
}

func TestConverToJT(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"ALTER TABLE `GiftBox` ADD COLUMN `isPublished` INT NOT NULL DEFAULT 0;", "--alter \"ADD COLUMN `isPublished` INT NOT NULL DEFAULT 0\" t=GiftBox"},
		{"ALTER TABLE User DROP COLUMN totalRevenueEarned;", "--alter \"DROP COLUMN totalRevenueEarned\" t=User"},
	}
	for _, test := range tests {
		b, q := ParseAlterQuery(test.input)
		if b {
			if got := ConverToJT(q); got != test.want {
				t.Error(got)
			}
		}
	}
}
