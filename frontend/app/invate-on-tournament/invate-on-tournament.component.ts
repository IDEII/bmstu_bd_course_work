import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { TournamentService } from '../tournament.service';
import { FighterService } from '../fighter.service';
import { Fighter } from '../fighter';
import { InvitedParticipant } from '../request';
import { RequestManagerService } from '../request-manager.service';

@Component({
  selector: 'app-invate-on-tournament',
  templateUrl: './invate-on-tournament.component.html',
  styleUrls: ['./invate-on-tournament.component.css'] 
})
export class InvateOnTournamentComponent implements OnInit {
  @Input() tournamentid: string = '';
  fighters: Fighter[] = [];
  invites: InvitedParticipant[] = [];
  inviteForm: FormGroup;
  deleteForm: FormGroup;
  isInvite: boolean = true;
  
  constructor(
    private formBuilder: FormBuilder,
    private tournamentService: TournamentService,
    private fighterService: FighterService,
    private requestManagerService: RequestManagerService
  ) {
    this.inviteForm = this.formBuilder.group({
      selectedFighter: ['']
    });
    this.deleteForm = this.formBuilder.group({
      selectedInvite: ['']
    });
  }
  
  ngOnInit(): void {
    this.loadInvites(); 

    this.fighterService.getFighters().subscribe((data: Fighter[]) => {
      this.fighters = data.filter(fighter => fighter.id !== '00000000-0000-0000-0000-000000000000');
    });
    this.inviteForm.get('selectedFighter')?.setValue(this.fighters.length > 0 ? this.fighters[0].id : '');
  }

  loadInvites(): void {
    this.requestManagerService.getInvited(this.tournamentid).subscribe((data: InvitedParticipant[]) => {
      this.invites = data;
      console.log(JSON.stringify(this.invites))
      this.deleteForm.get('selectedInvite')?.setValue(this.invites.length > 0 ? this.invites[0].InvitedParticipant_id : '');
    });
  }

  getFighterName(fighter_in_inviteId: string): string {
    const participant = this.fighters.find(fighter => fighter.id === fighter_in_inviteId);
    return participant ? participant.name : 'TBD';
  }

  inviteFighter(): void {
    const selectedFighterId = this.inviteForm.get('selectedFighter')?.value;
    if (this.tournamentid && selectedFighterId) {
      this.tournamentService.inviteRequest(this.tournamentid, selectedFighterId).subscribe({
        next: (response: any) => {
          console.log("Файтер приглашен", response);
          this.loadInvites(); 
        },
        error: (err: any) => console.error("Ошибка при приглашении", err)
      });
    } else {
      console.error("Ошибка: Необходимо выбрать бойца и турнир");
    }
  }

  deleteInvite(): void {
    const selectedInviteId = this.deleteForm.get('selectedInvite')?.value;
    console.log(selectedInviteId)
    if (selectedInviteId) {
      this.tournamentService.deleteInvite(this.tournamentid, selectedInviteId).subscribe({
        next: (response: any) => {
          console.log("Приглашение удалено", response);
          this.loadInvites();
        },
        error: (err: any) => console.error("Ошибка при удалении приглашения", err)
      });
    } else {
      console.error("Ошибка: Необходимо выбрать приглашение");
    }
  }
}