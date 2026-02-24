import { Component, Input } from '@angular/core';
import { OrganazerService } from '../organazer.service'
import { Organazer } from '../organazer';
import { ActivatedRoute, Router } from '@angular/router';
import { Tournament } from '../tournament'
import { TournamentService } from '../tournament.service';
import { FormGroup, FormControl, FormBuilder } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { TournamentComponent } from '../tournament/tournament.component';
import { UserService } from '../user.service';

@Component({
  selector: 'app-organazer-profile',
  templateUrl: './organazer-profile.component.html',
  styleUrl: './organazer-profile.component.css'
})
export class OrganazerProfileComponent {
  @Input() organazerID : string = ""
  editable : boolean = false;
  organazer : Organazer = {
    user_id: "",
    id: "",
    title: "",
    description: "",
    address: "",
    contact: "",
    founded_date: new Date()
 };
 tournament : Tournament = {
  id: '',
  name: '',
  address: '',
  startDate: new Date(),
  endDate: new Date(),
  organazer: '',
  rounds: '',
  category: 0,
};
showMng:boolean = false;
showCrt: boolean = false;
requestForm: FormGroup;
inputId : string = ""
app_tournament = TournamentComponent;
showInv: boolean = false;
  showTournament: boolean = false;
 org_id : string = "";
 editingTournamentId: string | null = null;
 tournaments: Array<Tournament>;
 selectedTournamentId: string | null = null;
 isEditing: boolean = false;
 isEditingTournament: boolean = false;
  showConfirmDelete: boolean = false; 
  form: FormGroup = new FormGroup ({
    id: new FormControl(''),
    title: new FormControl(''),
    description: new FormControl(''),
    address: new FormControl(''),
    contact: new FormControl(''),
    founded_date: new FormControl(new Date()),
  })
  formTournament: FormGroup = new FormGroup ({
    id: new FormControl(''),
    name: new FormControl(''),
    address: new FormControl(''),
    organazer: new FormControl(''),
    startDate: new FormControl(new Date()),
    endDate: new FormControl(new Date()),
    rounds: new FormControl(0),
    category: new FormControl(0),
  })

 constructor(private datepipe: DatePipe, private formBuilder: FormBuilder, private route: ActivatedRoute, private organazerService: OrganazerService, private tournamentService: TournamentService, private router: Router, private userService: UserService) { 
  this.tournaments = new Array<Tournament>();
  this.requestForm = this.formBuilder.group({
    selectedTournamentId: ['']
  });
 }
  ngOnInit() {
    this.form = this.formBuilder.group({
      id: ["",],
      title: ["",],
      description: ["",],
      address: ["",],
      contact: ["",],
      founded_date: [new Date().getDate(),],
    },
    );

    this.formTournament = this.formBuilder.group({
      id: ["",],
      name: ["", ],
      address: ['',],
      organazer: ["",],
      startDate: [new Date(),],
      endDate: [new Date(),],
      rounds: [0,],
      category: [0,],
    },
    );
    var id : string | null = this.route.snapshot.paramMap.get('id');
    if (this.organazerID !== "") {
        id = this.organazerID
        this.editable = true;
    } 
    if (id !== null) {
    this.org_id = id;

    this.organazerService.getOrganazer(id).subscribe(data => {
      this.organazer = data;
    });
    this.organazerService.getOrgTournaments(id).subscribe(data => {
      this.tournaments = data;
   });
  }
  }
  selectTournament(tournamentId: string) {
    this.selectedTournamentId = tournamentId;
  }

  editingTournament(tournamentId: string) {
    this.editingTournamentId = tournamentId;
    this.isEditingTournament = true;
  }

  showTour(tournamentId: string) {
    this.showTournament = !this.showTournament
    this.inputId = tournamentId
    console.log(this.inputId)
  }
  showInvite(tournamentId: string) {
    this.inputId = tournamentId
    this.showInv = !this.showInv
  }
  showCreate() {
    this.showCrt = !this.showCrt
  }

  showManage() {
    this.showMng = !this.showMng
  }

  updOrganazer() {
    const tempDate = this.datepipe.transform(this.organazer.founded_date, 'yyyy-MM-dd')
    const formData = {
        id: this.organazer.id,
        founded_date: tempDate
      }
    
    this.form.patchValue(formData)
    console.log(this.form.value)

    const id = this.organazer.id || this.route.snapshot.paramMap.get('id');
    if (id !== null) {
    this.organazerService.updateOrganazer(id, this.form.value).subscribe(response => {
            console.log('good thing');
            this.isEditing = false;
        }, error => {
          console.log('bad thing', error)
        }
    );
    
      this.organazerService.getOrganazer(id).subscribe(data => {
        this.organazer = data;
      });
      this.organazerService.getOrgTournaments(id).subscribe(data => {
        this.tournaments = data;
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

  cancelEditTournament() {
    this.isEditingTournament = false;
    this.editingTournamentId = null;
    this.formTournament = this.formBuilder.group({
        id: ["",],
        name: ["", ],
        address: ['',],
        organazer: ["",],
        startDate: [new Date(),],
        endDate: [new Date(),],
        rounds: [0,],
        category: [0,],
      },
      );
  }

  updTournament(tournament_id: string, rounds: string, category: number) {
    const dataPatch = {
      id: tournament_id,
      rounds: Number(rounds),
      category: category
    };
    this.formTournament.patchValue(dataPatch);
    console.log(JSON.stringify(this.formTournament.value));
    this.tournamentService.updateTournament(tournament_id, this.formTournament.value).subscribe(response => {
      console.log('good thing');
      this.cancelEditTournament();
      this.organazerService.getOrgTournaments(this.org_id).subscribe(data => {
        this.tournaments = data;
     });
    }, error => {
      console.log("bad thing", error);
      this.cancelEditTournament();
    });
  }

  delOrganazer() {
    this.organazerService.deleteOrganazer(this.organazer.id).subscribe(() => {
      console.log('организатор успешно удален');
      this.router.navigate(['/organazers']);
    }, error => {
      console.error('Ошибка при удалении организатора', error);
    });
  }


  conductTournament(tournamentId: string) {
    this.tournamentService.conductTournament(tournamentId).subscribe(
      response => {
        console.log('Tournament conducted successfully:', response);
      },
      error => {
        console.error('Error conducting tournament:', error);
      }
    );
  }

  delTournament(tournamentId: string) {
    this.tournamentService.deleteTournament(tournamentId).subscribe({
      next: (response: any) => {
        console.log("Удалено", response);
        this.organazerService.getOrgTournaments(this.org_id).subscribe(data => {
          this.tournaments = data;
       });
      },
      error: (err: any) => console.error("Ошибка при удалении турнира", err)
    })
  }
}