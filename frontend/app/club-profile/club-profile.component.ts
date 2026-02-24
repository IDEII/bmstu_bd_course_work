import { Component, Input, OnInit, resolveForwardRef, TemplateRef, ViewChild } from '@angular/core';
import { ClubMember } from '../club-member'
import { NgTemplateOutlet } from '@angular/common';
import { HttpClientModule } from '@angular/common/http'
import { ClubService } from '../club.service';
import { Router } from '@angular/router'
import { ActivatedRoute } from '@angular/router';
import { Club } from '../club'
import { Fighter } from '../fighter';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { UserService } from '../user.service';

@Component({
  selector: 'app-club-profile',
  templateUrl: './club-profile.component.html',
  styleUrls: ['./club-profile.component.css']
})

export class ClubProfileComponent implements OnInit {
    @Input() clubId: string = "";  
    editable : boolean = false;
    club : Club = {
      user_id: '0',
      id: '0',
      title: '',
      description: '',
      address: '',
      contact: '',
      founded_date: new Date(),
      rating: ''
  };
  fighters: Fighter[] = [];
  selectedFighterId: string | null;
  showAddMember: boolean = false;
  user_id : string = "";
  members: Array<Fighter> 
  isEditing: boolean = false;
  showConfirmDelete: boolean = false;
  form: FormGroup = new FormGroup ({
    id: new FormControl(''),
    title: new FormControl(''),
    description: new FormControl(''),
    address: new FormControl(''),
    contact: new FormControl(''),
    rating: new FormControl(''),
    founded_date: new FormControl(new Date()),
  })
  constructor(
    private datepipe: DatePipe,
    private formBuilder: FormBuilder,
    private route: ActivatedRoute,
    private router: Router,
    private clubService: ClubService, 
    private userService: UserService
  ) {
    this.members = new Array<Fighter>();
    this.selectedFighterId = "";
  }

  ngOnInit() {
    this.user_id = this.userService.get_user_id();

    this.form = this.formBuilder.group({
        title: ["",
  
        ],
        description: ["",
  
        ],
        rating: ["",
  
        ],
        address: ["",
  
        ],
        contact: ["",
  
        ],
        founded_date: [new Date().getDate(),
  
        ],
        id: ["",],
      },
      );
      const placeholder = {
        title: "1",
        description: "1",
        address: "1",
        contact: "1",
        rating: "1"
      } 
      this.form.patchValue(placeholder)
    var id : string | null = "";
    if (this.clubId !== "") {
      id = this.clubId;
      this.editable = true;
    } else {
      id = this.route.snapshot.paramMap.get('id');
    }
    console.log(id);
    if (id !== null) {
    this.clubService.getClub(id).subscribe(data => {
      this.club = data;
    });
    this.clubService.getMembers(id).subscribe(data => {
       this.members = data;
    });
  }
  this.loadFighterWithOutClub()
  }

  updClub() {
    
    const tempDate = this.datepipe.transform(this.club.founded_date, 'yyyy-MM-dd')
    const formData = {
        id: this.club.id,
        founded_date: tempDate
      }
    
    this.form.patchValue(formData)
    console.log(this.form.value)

    const id = this.clubId || this.route.snapshot.paramMap.get('id');
    if (id !== null) {
    this.clubService.updateClub(id, this.form.value).subscribe(response => {
            console.log('good thing');
            this.isEditing = false;
        }, error => {
          console.log('bad thing', error)
        }
    );
    
    this.clubService.getClub(id).subscribe(data => {
      this.club = data;
    });
    
    this.clubService.getMembers(id).subscribe(data => {
        this.members = data;
    });
}
  }
  cancelEdit() {
    this.isEditing = false;
    this.form = this.formBuilder.group({
        title: ["",
  
        ],
        description: ["",
  
        ],
        rating: ["",
  
        ],
        address: ["",
  
        ],
        contact: ["",
  
        ],
        founded_date: [new Date(),
  
        ],
        id: ["",],
      },
      );
  }

  loadFighterWithOutClub() {
    this.clubService.getFightersWithOutClub().subscribe(data => {
      this.fighters = data.filter(fighter => fighter.id !== '00000000-0000-0000-0000-000000000000');
   });
  }

  kickMember(member_id: string) {
    this.clubService.kickFromClub(this.club.id, member_id).subscribe(response => {
      console.log(response);
    });
    this.clubService.getMembers(this.clubId).subscribe(data => {
      this.members = data;
  });
  }

  addMember() {
    if (this.selectedFighterId) {
      this.clubService.addMemberToClub(this.club.id, this.selectedFighterId).subscribe(response => {
        console.log(response) 
        this.showAddMember = false;
        this.selectedFighterId = null; 
      }, error => {
        console.error('Ошибка при добавлении участника', error);
      });
    }
  }

  deleteClub() {
    this.clubService.deleteClub(this.club.id).subscribe(() => {
      console.log('Клуб успешно удален');
      this.router.navigate(['/clubs']);

    }, error => {
      console.error('Ошибка при удалении клуба', error);
    });
  }
}
  