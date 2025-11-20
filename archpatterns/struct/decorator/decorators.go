package main

type SQLSkillDeveloperDecorator struct {
	Developer
}

func NewSQLSkillDecorator(developer Developer) Developer {
	return &SQLSkillDeveloperDecorator{Developer: developer}
}

// Skills добавляет новый скилл в другой участок памяти чтобы не менять исходный слайс со сикллами
func (s *SQLSkillDeveloperDecorator) Skills() []string {
	developerSkills := s.Developer.Skills()
	newSkills := make([]string, len(developerSkills), len(developerSkills)+1)

	copy(newSkills, developerSkills)

	return append(newSkills, "SQL")
}
