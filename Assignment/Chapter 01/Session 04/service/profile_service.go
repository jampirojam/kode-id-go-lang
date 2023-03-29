package service

import (
	"fmt"
	"ojam-test/fourth-session/model"
)

func AddNewProfile(profiles []model.Profile) []model.Profile {
	fmt.Println("INPUT NEW PROFILE")

	inputName := InputString("Name")
	inputAge := InputInteger("Age")
	inputJob := InputString("Job")
	inputAddress := InputString("Address")
	inputMarried := InputCharacter("Married [S/M/W/D]")

	statusMarried, isMarried := GetMarriedStatus(inputMarried)

	fmt.Println(statusMarried, isMarried)
	
	var id int
	for i := 0; i < len(profiles); i++ {
		if i == (len(profiles) - 1) {
			id = profiles[i].Id + 1
		}
	}

	if id == 0 {
		id = 1
	}

	profile := model.Profile{}
	profile.Id = id
	profile.Name = inputName
	profile.Age = inputAge
	profile.Job = inputJob
	profile.Address = inputAddress
	profile.Married = isMarried
	profile.StatusMarried = statusMarried

	profiles = append(profiles, profile)
	SuccessMessage(fmt.Sprintf("Sucessfully add new profile with Id: %v",  profile.Id))
	fmt.Printf("\n\n")

	return profiles
}

func DeleteProfileById(profiles []model.Profile) []model.Profile {
	fmt.Println("DELETE A PROFILE")

	GetProfiles(profiles)

	inputId := InputInteger("Profile Id")

	lenght := len(profiles)
	idList := make([]int, lenght)
	for i := 0; i < len(profiles); i++ {
		idList[i] = profiles[i].Id
	}

	var idFound bool
	for i := 0; i < lenght; i++ {
		if idList[i] == inputId {
			idFound = true
		}
	}

	newProfiles := []model.Profile{}
	if idFound {
		for i := 0; i < lenght; i++ {
			if inputId == profiles[i].Id {
				SuccessMessage(fmt.Sprintf("Profile with Id: %v, successfully deleted", profiles[i].Id))
			} else {
				newProfiles = append(newProfiles, profiles[i])
			}
		}

		return newProfiles
	} else {
		ErrorMessage(fmt.Sprintf("Profile with Id: %d is not found", inputId))
	}
	fmt.Printf("\n\n")

	return profiles
}

func GetProfiles(profiles []model.Profile) []model.Profile {
	fmt.Println("LIST OF PROFILE")
	fmt.Printf("| %-5s | %-25s | %-5s | %-20s | %-25s | %-10s | %-15s |\n",
		"ID",
		"Name",
		"Age",
		"Job",
		"Address",
		"Married",
		"Marrital Status",
	)
	for _, p := range profiles {
		fmt.Printf("| %-5s | %-25s | %-5s | %-20s | %-25s | %-10s | %-15s |\n",
			"—————",
			"—————————————————————————",
			"—————",
			"————————————————————",
			"—————————————————————————",
			"——————————",
			"———————————————",
		)
		fmt.Printf("| %-5v | %-25v | %-5v | %-20v | %-25v | %-10v | %-15v |\n",
			p.Id,
			p.Name,
			p.Age,
			p.Job,
			p.Address,
			p.Married,
			p.StatusMarried,
		)
	}
	fmt.Printf("\n\n")

	return profiles
}

func UpdateProfileById(profiles []model.Profile) []model.Profile {
	fmt.Println("UPDATE A PROFILE")
	GetProfiles(profiles)

	inputId := InputInteger("Profile Id")

	lenght := len(profiles)
	idList := make([]int, lenght)
	for i := 0; i < len(profiles); i++ {
		idList[i] = profiles[i].Id
	}

	var idFound bool
	for i := 0; i < lenght; i++ {
		if idList[i] == inputId {
			idFound = true
		}
	}

	newProfiles := []model.Profile{}
	if idFound {
		for i := 0; i < lenght; i++ {
			if inputId == profiles[i].Id {
				profile := model.Profile{}

				profile.Id = profiles[i].Id

				input := InputYesOrNo("Update profile name [y/n]")
				if input {
					profile.Name = InputString("Name")
				} else {
					profile.Name = profiles[i].Name
				}

				input = InputYesOrNo("Update profile age [y/n]")
				if input {
					profile.Age = InputInteger("Age")
				} else {
					profile.Age = profiles[i].Age
				}

				input = InputYesOrNo("Update profile job [y/n]")
				if input {
					profile.Job = InputString("Job")
				} else {
					profile.Job = profiles[i].Job
				}

				input = InputYesOrNo("Update profile address [y/n]")
				if input {
					profile.Address = InputString("Address")
				} else {
					profile.Address = profiles[i].Address
				}

				input = InputYesOrNo("Update profile married [y/n]")
				if input {
					inputMarried := InputCharacter("Married [S/M/W/D]")
					statusMarried, isMarried := GetMarriedStatus(inputMarried)
					profile.Married = isMarried
					profile.StatusMarried = statusMarried
				} else {
					profile.Married = profiles[i].Married
					profile.StatusMarried = profiles[i].StatusMarried
				}

				newProfiles = append(newProfiles, profile)
				fmt.Printf("Profile with Id: %v has successfully updated\n\n", profiles[i].Id)
			} else {
				newProfiles = append(newProfiles, profiles[i])
			}
		}
		fmt.Println()

		return newProfiles
	} else {
		ErrorMessage(fmt.Sprintf("Profile with Id: %d is not found", inputId))
	}
	fmt.Printf("\n\n")

	return profiles
}

func SearchProfileById(profiles []model.Profile) []model.Profile {
	fmt.Println("SEARCH PROFILE")
	inputId := InputInteger("Profile Id")

	lenght := len(profiles)
	idList := make([]int, lenght)
	for i := 0; i < len(profiles); i++ {
		idList[i] = profiles[i].Id
	}

	var idFound bool
	for i := 0; i < lenght; i++ {
		if idList[i] == inputId {
			idFound = true
		}
	}

	if idFound {
		fmt.Printf("| %-5s | %-25s | %-5s | %-20s | %-25s | %-10s | %-15s |\n",
			"ID",
			"Name",
			"Age",
			"Job",
			"Address",
			"Married",
			"Marrital Status",
		)
		for i := 0; i < lenght; i++ {
			if inputId == profiles[i].Id {
				fmt.Printf("| %-5s | %-25s | %-5s | %-20s | %-25s | %-10s | %-15s |\n",
					"—————",
					"—————————————————————————",
					"—————",
					"————————————————————",
					"—————————————————————————",
					"——————————",
					"———————————————",
				)
				fmt.Printf("| %-5v | %-25v | %-5v | %-20v | %-25v | %-10v | %-15v |\n",
					profiles[i].Id,
					profiles[i].Name,
					profiles[i].Age,
					profiles[i].Job,
					profiles[i].Address,
					profiles[i].Married,
					profiles[i].StatusMarried,
				)
			}
		}
		fmt.Println()
	} else {
		ErrorMessage(fmt.Sprintf("Profile with Id: %d is not found", inputId))
	}
	fmt.Printf("\n\n")

	return profiles
}
