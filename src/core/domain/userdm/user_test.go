package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestUser_NewUser_Success(t *testing.T) {
	userId := userdm.NewUserID()
	userName, _ := userdm.NewUserName("Test User")
	password, _ := userdm.NewPassword("validPassword0")
	email, _ := userdm.NewEmail("test@example.com")
	selfIntroduction, _ := userdm.NewSelfIntroduction("よろしくお願いします")

	tagId, _ := tagdm.NewTagID("test-tag-id")
	skill, _ := userdm.NewSkill(userdm.NewSkillID(), tagId, 5, 3)
	skills := []userdm.Skill{*skill}

	careerDetail, _ := userdm.NewCareerDetail("Web開発に従事")
	careerStartYear, _ := userdm.NewCareerStartYear(2020)
	careerEndYear, _ := userdm.NewCareerEndYear(2022)
	career, _ := userdm.NewCareer(
		userdm.NewCareerID(),
		*careerDetail,
		*careerStartYear,
		*careerEndYear,
	)
	careers := []userdm.Career{*career}

	user, err := userdm.NewUser(userId, *userName, *password, *email, skills, careers, selfIntroduction)
	if err != nil {
		t.Errorf("NewUser() with valid parameters should not return error, got: %v", err)
	}

	if user == nil {
		t.Error("NewUser() should not return nil")
	}

	if user.ID() != userId {
		t.Error("Id() should return correct userId")
	}

	if user.Name() != *userName {
		t.Error("Name() should return correct userName")
	}

	if user.Password() != *password {
		t.Error("Password() should return correct password")
	}

	if user.Email() != *email {
		t.Error("Email() should return correct email")
	}

	if user.SelfIntroduction() != selfIntroduction {
		t.Error("SelfIntroduction() should return correct selfIntroduction")
	}

	if len(user.Skills()) != 1 {
		t.Error("Skills() should return correct number of skills")
	}

	if len(user.Careers()) != 1 {
		t.Error("Careers() should return correct number of careers")
	}
}

func TestUser_NewUser_EmptySkills(t *testing.T) {
	userId := userdm.NewUserID()
	userName, _ := userdm.NewUserName("Test User")
	password, _ := userdm.NewPassword("validPassword0")
	email, _ := userdm.NewEmail("test@example.com")
	selfIntroduction, _ := userdm.NewSelfIntroduction("よろしくお願いします")

	skills := []userdm.Skill{}
	careers := []userdm.Career{}

	_, err := userdm.NewUser(userId, *userName, *password, *email, skills, careers, selfIntroduction)
	if err == nil {
		t.Error("NewUser() with empty skills should return error")
	}
}
