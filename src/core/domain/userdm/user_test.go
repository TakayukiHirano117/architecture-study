package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestUser_NewUser_Success(t *testing.T) {
	userId := userdm.NewUserID()

	userName, err := userdm.NewUserName("Test User")
	if err != nil {
		t.Errorf("NewUserName() with valid parameters should not return error, got: %v", err)
	}

	password, err := userdm.NewPassword("validPassword0")
	if err != nil {
		t.Errorf("NewPassword() with valid parameters should not return error, got: %v", err)
	}

	email, err := userdm.NewEmail("test@example.com")
	if err != nil {
		t.Errorf("NewEmail() with valid parameters should not return error, got: %v", err)
	}

	selfIntroduction, err := userdm.NewSelfIntroduction("よろしくお願いします")
	if err != nil {
		t.Errorf("NewSelfIntroduction() with valid parameters should not return error, got: %v", err)
	}

	tagId := tagdm.NewTagID()

	tagName, err := tagdm.NewTagName("test-tag-name")
	if err != nil {
		t.Errorf("NewTagName() with valid parameters should not return error, got: %v", err)
	}

	tag, err := tagdm.NewTag(tagId, *tagName)
	if err != nil {
		t.Errorf("NewTag() with valid parameters should not return error, got: %v", err)
	}

	skill, err := userdm.NewSkill(userdm.NewSkillID(), tag, 5, 3)
	if err != nil {
		t.Errorf("NewSkill() with valid parameters should not return error, got: %v", err)
	}

	skills := []userdm.Skill{*skill}

	careerDetail, err := userdm.NewCareerDetail("Web開発に従事")
	if err != nil {
		t.Errorf("NewCareerDetail() with valid parameters should not return error, got: %v", err)
	}

	careerStartYear, err := userdm.NewCareerStartYear(2020)
	if err != nil {
		t.Errorf("NewCareerStartYear() with valid parameters should not return error, got: %v", err)
	}

	careerEndYear, err := userdm.NewCareerEndYear(2022)
	if err != nil {
		t.Errorf("NewCareerEndYear() with valid parameters should not return error, got: %v", err)
	}

	career, err := userdm.NewCareer(
		userdm.NewCareerID(),
		*careerDetail,
		*careerStartYear,
		*careerEndYear,
	)
	if err != nil {
		t.Errorf("NewCareer() with valid parameters should not return error, got: %v", err)
	}

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
	userName, err := userdm.NewUserName("Test User")
	if err != nil {
		t.Errorf("NewUserName() with valid parameters should not return error, got: %v", err)
	}
	password, err := userdm.NewPassword("validPassword0")
	if err != nil {
		t.Errorf("NewPassword() with valid parameters should not return error, got: %v", err)
	}
	email, err := userdm.NewEmail("test@example.com")
	if err != nil {
		t.Errorf("NewEmail() with valid parameters should not return error, got: %v", err)
	}
	selfIntroduction, err := userdm.NewSelfIntroduction("よろしくお願いします")
	if err != nil {
		t.Errorf("NewSelfIntroduction() with valid parameters should not return error, got: %v", err)
	}

	skills := []userdm.Skill{}
	careers := []userdm.Career{}

	_, err = userdm.NewUser(userId, *userName, *password, *email, skills, careers, selfIntroduction)
	if err == nil {
		t.Error("NewUser() with empty skills should return error")
	}
}
