package menu

import (
	"fmt"
	"os"
	"ojam-test/fourth-session/model"
	"ojam-test/fourth-session/service"
)

func ParentMenu() {
	profiles := []model.Profile{}
	menu(profiles)
}

func menu(profiles []model.Profile) {
	fmt.Println("ADMIN PROFILE MENU ============")
	fmt.Println("1. Add Profile")
	fmt.Println("2. Delete Profile By Id")
	fmt.Println("3. Update Profile By Id")
	fmt.Println("4. Get All Profiles")
	fmt.Println("5. Search Profile By Id")
	choice := service.InputInteger("Choose number")
	
	if choice == 1 {
		profiles = service.AddNewProfile(profiles)
		menu(profiles)
	} else if choice == 2 {
		profiles = service.DeleteProfileById(profiles)
		menu(profiles)
	} else if choice == 3 {
		profiles = service.UpdateProfileById(profiles)
		menu(profiles)
	} else if choice == 4 {
		profiles = service.GetProfiles(profiles)
		menu(profiles)
	} else if choice == 5 {
		profiles = service.SearchProfileById(profiles)
		menu(profiles)
	} else {
		os.Exit(0)
	}
}
