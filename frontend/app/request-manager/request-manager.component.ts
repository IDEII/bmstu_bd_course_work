import { Component, Input, OnInit } from '@angular/core';
import { RequestManagerService } from '../request-manager.service';
import { TournamentService } from '../tournament.service';
import { FighterService } from '../fighter.service';
import { Request, InvitedParticipant } from '../request';
import { Tournament } from '../tournament';
import { Fighter } from '../fighter';
import { OrganazerProfileComponent } from '../organazer-profile/organazer-profile.component';
import { OrganazerService } from '../organazer.service';

@Component({
  selector: 'app-request-manager',
  templateUrl: './request-manager.component.html',
  styleUrl: './request-manager.component.css'
})

export class RequestManagerComponent implements OnInit {
  @Input() organazer_id : string = ""
  requests: Request[] = [];
  tournaments: Tournament[] = [];
  selectedTournament: Tournament | null = null;
  InvitedParticipants: InvitedParticipant[] = [];
  tour_id: string = "";
  fighters: Fighter[] = [];
  showInv: boolean = false;

  showInvite() {
    this.showInv = !this.showInv
  }

  constructor(
    private requestService: RequestManagerService,
    private organazerService: OrganazerService,
    private tournamentService: TournamentService,
    private fighterService: FighterService
  ) {}
  
  ngOnInit(): void {
    this.organazerService.getOrgTournaments(this.organazer_id).subscribe({
      next: (data: any[]) => this.tournaments = data
    });

    this.fighterService.getFighters().subscribe({
      next: (data: any[]) => this.fighters = data
    });
  }
  
  onTournamentSelect(event: Event) {
    const selectElement = event.target as HTMLSelectElement;
    const selectedTournamentId = selectElement.value;
    this.tour_id = selectedTournamentId;
    this.selectedTournament = this.tournaments.find(t => t.id === selectedTournamentId) || null;
    this.requests = [];
    this.InvitedParticipants = [];
    this.loadRequestsAndInvited(this.tour_id)
  }

  loadRequestsAndInvited(tournamentId: string): void {
    this.requestService.getRequests(tournamentId).subscribe({next: (data: Request[]) => this.requests = data});
    console.log(JSON.stringify(this.requests))
    this.requestService.getInvited(tournamentId).subscribe({next: (data: InvitedParticipant[]) => this.InvitedParticipants = data});
  }
  

  approveRequest(requestId: string): void {
    this.requestService.approveRequest(requestId).subscribe(() => {
      this.loadRequestsAndInvited(this.tour_id);
    });
  }

  denyRequest(requestId: string): void {
    this.requestService.denyRequest(requestId).subscribe(() => {
      this.loadRequestsAndInvited(this.tour_id);
    });
  }

  getFighterName(fighterId: string): string | undefined {
    const fighter = this.fighters.find(f => f.id === fighterId);
    return fighter ? fighter.name : undefined;
  }
}